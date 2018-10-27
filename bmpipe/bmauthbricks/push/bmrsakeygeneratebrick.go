package authpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmsecurity"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BMRsaKeyGenerateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMRsaKeyGenerateBrick) Exec() error {
	rsaKey, err := bmsecurity.GetRsaKey(512)
	if err != nil {
		return err
	}
	tmp := b.bk.Pr.(auth.BMRsaKey)
	tmp.PublicKey = rsaKey.PublicKey
	tmp.PrivateKey = rsaKey.PrivateKey
	tmp.InsertBMObject()
	tmp.PrivateKey = ""
	b.BrickInstance().Pr = tmp
	return nil
}

func (b *BMRsaKeyGenerateBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMRsaKey)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMRsaKeyGenerateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMRsaKeyGenerateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMRsaKeyGenerateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMRsaKey)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMRsaKeyGenerateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMRsaKey = b.BrickInstance().Pr.(auth.BMRsaKey)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
