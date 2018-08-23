package coursepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/course"
	"io"
	"net/http"
)

type BMCoursePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMCoursePushBrick) Exec() error {
	var tmp course.BMCourse = b.bk.Pr.(course.BMCourse)
	if tmp.Id != "" && tmp.Id_.Valid() {
		if tmp.Valid() && tmp.IsRegistered() {
			//b.bk.Err = -3
			b.bk.Pr = tmp
		} else {
			tmp.InsertBMObject()
			b.bk.Pr = tmp
		}
	}

	return nil
}

func (b *BMCoursePushBrick) Prepare(pr interface{}) error {
	req := pr.(course.BMCourse)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMCoursePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMCoursePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMCoursePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(course.BMCourse)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMCoursePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval course.BMCourse = b.BrickInstance().Pr.(course.BMCourse)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
