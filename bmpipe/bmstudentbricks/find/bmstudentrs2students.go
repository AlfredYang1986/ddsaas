package studentfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/student"
	"io"
	"net/http"
)

type BMStudentRS2StudentsBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudentRS2StudentsBrick) Exec() error {
	props := b.bk.Pr.(student.BMStudentsProp)
	var revals student.BMStudents
	for _, prop := range props.StudentsProp {
		reval, err := findStudent(prop)
		guard, err := findGuardians(prop)
		if err != nil {
			return err
		}
		reval.Guardians = guard
		revals.Students = append(revals.Students, reval)
	}
	b.bk.Pr = revals
	return nil
}

func (b *BMStudentRS2StudentsBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudentsProp)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudentRS2StudentsBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudentRS2StudentsBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudentRS2StudentsBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudents)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudentRS2StudentsBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudents = b.BrickInstance().Pr.(student.BMStudents)
		jsonapi.ToJsonAPI(reval.Students, w)
	}
}
