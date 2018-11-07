package teacherpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
)

type BmTeacherPersonPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherPersonPushBrick) Exec() error {
	var tmp = b.bk.Pr.(teacher.BmTeacher)
	ap := tmp.Person
	ap.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BmTeacherPersonPushBrick) Prepare(pr interface{}) error {
	req := pr.(teacher.BmTeacher)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTeacherPersonPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTeacherPersonPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherPersonPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BmTeacher)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTeacherPersonPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeacher = b.BrickInstance().Pr.(teacher.BmTeacher)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
