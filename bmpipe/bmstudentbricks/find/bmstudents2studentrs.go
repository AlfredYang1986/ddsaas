package studentfind

import (
	"fmt"
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

type BMStudents2StudentRSBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudents2StudentRSBrick) Exec() error {
	var tmps student.BMStudents = b.bk.Pr.(student.BMStudents)
	var revals student.BMStudentsProp

	for _, v := range tmps.Students {
		var tmp student.BMStudent = v
		eq := request.EQCond{}
		eq.Ky = "student_id"
		eq.Vy = tmp.Id
		req := request.Request{}
		req.Res = "BMStudentProp"
		var condi []interface{}
		condi = append(condi, eq)
		c := req.SetConnect("conditions", condi)
		fmt.Println(c)

		var reval student.BMStudentProp
		err := reval.FindOne(c.(request.Request))
		if err != nil {
			return err
		}
		revals.StudentsProp = append(revals.StudentsProp, reval)
	}

	b.bk.Pr = revals
	return nil
}

func (b *BMStudents2StudentRSBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudents)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudents2StudentRSBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudents2StudentRSBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudents2StudentRSBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudentsProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudents2StudentRSBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudents = b.BrickInstance().Pr.(student.BMStudents)
		jsonapi.ToJsonAPI(reval.Students, w)
	}
}
