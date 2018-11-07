package teacherfind

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
)

type BmTeacherMultiRS struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherMultiRS) Exec() error {
	var tmp teacher.BmTeachers = b.bk.Pr.(teacher.BmTeachers)
	var err error

	req1 := request.Request{}
	req1.Res = "BMTeacherProp"
	var condi1 []interface{}
	for _, item := range tmp.Teachers {
		eq := request.Eqcond{}
		eq.Ky = "teacherId"
		eq.Vy = item.Id
		condi1 = append(condi1, eq)
	}

	c1 := req1.SetConnect("conditions", condi1)
	var tp teacher.BmTeacherProps
	err = tp.FindMulti(c1.(request.Request))
	if err != nil {
		return err
	}

	for _, item := range tmp.Teachers {
		for i, p := range tp.TeacherProps {
			if p.TeacherId == item.Id {
				_, per := p.GetPerson()
				rel := item.SetConnect("person", per)
				tmp.Teachers[i] = rel.(teacher.BmTeacher)
			}
		}
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmTeacherMultiRS) Prepare(pr interface{}) error {
	req := pr.(teacher.BmTeachers)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTeacherMultiRS) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTeacherMultiRS) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherMultiRS) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BmTeacher)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTeacherMultiRS) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeachers = b.BrickInstance().Pr.(teacher.BmTeachers)
		jsonapi.ToJsonAPI(&reval, w)
	}
}