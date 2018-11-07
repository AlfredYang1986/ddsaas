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

type BmTeacherRS2Teacher struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherRS2Teacher) Exec() error {
	var tmp teacher.BmTeacher = b.bk.Pr.(teacher.BmTeacher)
	var err error

	eq := request.Eqcond{}
	eq.Ky = "teacherId"
	eq.Vy = tmp.Id
	req1 := request.Request{}
	req1.Res = "BMTeacherProp"
	var condi1 []interface{}
	condi1 = append(condi1, eq)
	c1 := req1.SetConnect("conditions", condi1)
	var teacherProp teacher.BMTeacherProp
	err = teacherProp.FindOne(c1.(request.Request))
	if err != nil {
		return err
	}

	err, person := teacherProp.GetPerson()
	tmp.Person = person

	b.bk.Pr = tmp
	return err
}

func (b *BmTeacherRS2Teacher) Prepare(pr interface{}) error {
	req := pr.(teacher.BmTeacher)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTeacherRS2Teacher) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTeacherRS2Teacher) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherRS2Teacher) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BmTeacher)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTeacherRS2Teacher) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeacher = b.BrickInstance().Pr.(teacher.BmTeacher)
		jsonapi.ToJsonAPI(&reval, w)
	}
}