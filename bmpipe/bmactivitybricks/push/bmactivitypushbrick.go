package activitypush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/activity"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BMActivityPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMActivityPushBrick) Exec() error {
	var tmp activity.BMActivity = b.bk.Pr.(activity.BMActivity)
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

func (b *BMActivityPushBrick) Prepare(pr interface{}) error {
	req := pr.(activity.BMActivity)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMActivityPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMActivityPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMActivityPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(activity.BMActivity)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMActivityPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval activity.BMActivity = b.BrickInstance().Pr.(activity.BMActivity)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
