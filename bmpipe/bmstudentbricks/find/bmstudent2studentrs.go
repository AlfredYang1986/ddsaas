package studentfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/student"
	"io"
	"net/http"
)

type BMStudent2StudentRSBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudent2StudentRSBrick) Exec() error {
	var tmp student.BMStudent = b.bk.Pr.(student.BMStudent)

	eq := request.EQCond{}
	eq.Ky = "student_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BMStudentProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	var reval student.BMStudentProp
	err := reval.FindOne(c.(request.Request))

	b.bk.Pr = reval
	return err
}

func (b *BMStudent2StudentRSBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudent)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudent2StudentRSBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudent2StudentRSBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudent2StudentRSBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudentProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudent2StudentRSBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudent = b.BrickInstance().Pr.(student.BMStudent)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
