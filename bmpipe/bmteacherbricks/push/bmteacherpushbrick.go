package teacherpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"io"
	"net/http"
	"time"
)

type BMTeacherPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMTeacherPushBrick) Exec() error {
	var tmp teacher.BMTeacher = b.bk.Pr.(teacher.BMTeacher)
	//TODO： use type Timestamp
	ts := time.Now().Unix()
	tmp.Found = ts
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BMTeacherPushBrick) Prepare(pr interface{}) error {
	req := pr.(teacher.BMTeacher)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMTeacherPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMTeacherPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMTeacherPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BMTeacher)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMTeacherPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BMTeacher = b.BrickInstance().Pr.(teacher.BMTeacher)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
