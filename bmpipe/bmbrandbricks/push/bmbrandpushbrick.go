package brandpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BmBrandPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBrandPushBrick) Exec() error {
	tmp := b.bk.Pr.(brand.BmBrand)

	var err error
	if tmp.Valid() && tmp.Id != "" && tmp.Id_.Valid() {
		//if tmp.IsBrandRegistered() {
		//	b.bk.Err = -5
		//}
		err = tmp.InsertBMObject()
	}
	b.bk.Pr = tmp
	return err
}

func (b *BmBrandPushBrick) Prepare(pr interface{}) error {
	req := pr.(brand.BmBrand)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBrandPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBrandPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBrandPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(brand.BmBrand)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBrandPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval brand.BmBrand = b.BrickInstance().Pr.(brand.BmBrand)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
