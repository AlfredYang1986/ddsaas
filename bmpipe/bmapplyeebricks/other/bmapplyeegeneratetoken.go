package applyeeother

import (
	"crypto/md5"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"io"
	"net/http"
)

type BmApplyeeGenerateToken struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyeeGenerateToken) Exec() error {
	var err error
	tmp := b.bk.Pr.(applyee.BmApplyee)

	h := md5.New()
	io.WriteString(h, tmp.Id)
	token := fmt.Sprintf("%x", h.Sum(nil))
	err = bmoauth.PushToken(token)

	bmls := auth.BmLoginSucceed{
		Id: tmp.Id,
		Id_: tmp.Id_,
		Applyee:tmp,
		Token:token,
	}

	b.BrickInstance().Pr = bmls
	return err
}

func (b *BmApplyeeGenerateToken) Prepare(pr interface{}) error {
	req := pr.(applyee.BmApplyee)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyeeGenerateToken) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	ec := b.BrickInstance().Err
	if int(idx) < tmp-1 && ec == 0 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyeeGenerateToken) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyeeGenerateToken) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(applyee.BmApplyee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyeeGenerateToken) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(auth.BmLoginSucceed)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

