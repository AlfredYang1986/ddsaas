package teather_person

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
)

type BmPersonTeacherRS struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmPersonTeacherRS) Exec() error {
	var tmp person.BmPersons = b.bk.Pr.(person.BmPersons)
	var err error

	var condi1 []interface{}
	req1 := request.Request{}
	for _, item := range tmp.Persons {
		eq := request.Eqcond{}
		eq.Ky = "personId"
		eq.Vy = item.Id
		condi1 = append(condi1, eq)
	}

	req1.Res = "BMTeacherProp"
	c1 := req1.SetConnect("conditions", condi1)
	var tp teacher.BmTeacherProps
	err = tp.FindMulti(c1.(request.Request))
	if err != nil {
		return err
	}

	var condi2 []interface{}
	req2 := request.Request{}
	for _, item := range tp.TeacherProps {
		eq := request.Eqcond{}
		eq.Ky = "id"
		eq.Vy = item.TeacherId
		condi2 = append(condi2, eq)
	}

	req2.Res = "BmTeacher"
	c2 := req2.SetConnect("conditions", condi2)
	var ths teacher.BmTeachers
	err = ths.FindMulti(c2.(request.Request))
	if err != nil {
		return err
	}

	var res []teacher.BmTeacher
	for _, item := range ths.Teachers {
		var per person.BmPerson
		for _, tp := range tp.TeacherProps {
			var pid string
			if tp.TeacherId == item.Id {
				pid = tp.PersonId
			}

			for _, p := range tmp.Persons {
				if p.Id == pid {
					per = p
				}
			}
		}
		res = append(res, item.SetConnect("person", per).(teacher.BmTeacher))
	}

	result := teacher.BmTeachers{}
	result.Teachers = res
	b.bk.Pr = result
	return err
}

func (b *BmPersonTeacherRS) Prepare(pr interface{}) error {
	req := pr.(person.BmPersons)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmPersonTeacherRS) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmPersonTeacherRS) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmPersonTeacherRS) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BmTeachers)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmPersonTeacherRS) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeachers = b.BrickInstance().Pr.(teacher.BmTeachers)
		jsonapi.ToJsonAPI(&reval, w)
	}
}