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

type BMPhonePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMPhonePushBrick) Exec() error {
	var tmp auth.BmAuth = b.bk.Pr.(auth.BmAuth)
	ap := tmp.Phone
	if ap.Id != "" && ap.Id_.Valid() {
		if ap.Valid() && ap.IsPhoneRegisted() {
			b.bk.Err = -1
		} else {
			ap.InsertBMObject()
		}
	}
	return nil
}

func (b *BMPhonePushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BmAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMPhonePushBrick) Done(pkg string, idx int64, e error) error {
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

func (b *BMPhonePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMPhonePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BmAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMPhonePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmAuth = b.BrickInstance().Pr.(auth.BmAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
