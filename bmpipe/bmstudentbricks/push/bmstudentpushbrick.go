package studentpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/student"
	"io"
	"net/http"
	"time"
)

type BMStudentPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudentPushBrick) Exec() error {
	var tmp student.BMStudent = b.bk.Pr.(student.BMStudent)
	//TODO： use type Timestamp
	ts := time.Now().Unix()
	tmp.Found = ts
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BMStudentPushBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudent)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudentPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudentPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudentPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudent)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudentPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudent = b.BrickInstance().Pr.(student.BMStudent)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
