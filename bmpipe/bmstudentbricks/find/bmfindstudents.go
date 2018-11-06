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

type BMStudentFindMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudentFindMultiBrick) Exec() error {
	var tmps student.BMStudents
	err := tmps.FindMulti(*b.bk.Req)
	var revals student.BMStudents
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

		var prop student.BMStudentProp
		err := prop.FindOne(c.(request.Request))
		if err != nil {
			return err
		}

		reval, err := findStudent(prop)
		guard, err := findGuardians(prop)
		if err != nil {
			return err
		}
		reval.Guardians = guard
		revals.Students = append(revals.Students, reval)

	}

	b.bk.Pr = revals
	return err
}

func (b *BMStudentFindMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BMStudentFindMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudentFindMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudentFindMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudents)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudentFindMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudents = b.BrickInstance().Pr.(student.BMStudents)
		jsonapi.ToJsonAPI(reval.Students, w)
	}
}
