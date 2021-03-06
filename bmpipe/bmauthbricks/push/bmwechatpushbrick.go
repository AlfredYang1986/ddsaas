package authpush

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"io"
	"net/http"
)

type BMWechatPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMWechatPushBrick) Exec() error {
	var tmp auth.BmAuth = b.bk.Pr.(auth.BmAuth)
	aw := tmp.Wechat
	if aw.Id != "" && aw.Id_.Valid() {
		if aw.Valid() && aw.IsWechatRegisted() {
			b.bk.Err = -2
		} else {
			aw.InsertBMObject()
		}
	}
	return nil
}

func (b *BMWechatPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BmAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMWechatPushBrick) Done(pkg string, idx int64, e error) error {
	ec := b.BrickInstance().Err
	if ec != 0 {
		return nil
	}
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMWechatPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMWechatPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BmAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMWechatPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmAuth = b.BrickInstance().Pr.(auth.BmAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
