package accountpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"io"
	"net/http"
)

type BmAccountBindBrand struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmAccountBindBrand) Exec() error {
	tmp := b.bk.Pr.(account.BmBindAccountBrand)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmAccountBindBrand) Prepare(pr interface{}) error {
	req := pr.(account.BmBindAccountBrand)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmAccountBindBrand) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmAccountBindBrand) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmAccountBindBrand) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(account.BmBindAccountBrand)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmAccountBindBrand) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval account.BmBindAccountBrand = b.BrickInstance().Pr.(account.BmBindAccountBrand)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

