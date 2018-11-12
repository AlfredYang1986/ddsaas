package tagimgpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"io"
	"net/http"
)

type BmTagImgBindYard struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTagImgBindYard) Exec() error {
	tmp := b.bk.Pr.(yard.BmBindYardImg)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmTagImgBindYard) Prepare(pr interface{}) error {
	req := pr.(yard.BmBindYardImg)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTagImgBindYard) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTagImgBindYard) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTagImgBindYard) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(yard.BmBindYardImg)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTagImgBindYard) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval yard.BmBindYardImg = b.BrickInstance().Pr.(yard.BmBindYardImg)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

