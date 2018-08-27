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

type BMEntitynamePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMEntitynamePushBrick) Exec() error {
	var tmp brand.BMBrand = b.bk.Pr.(brand.BMBrand)
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BMEntitynamePushBrick) Prepare(pr interface{}) error {
	req := pr.(brand.BMBrand)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMEntitynamePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMEntitynamePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMEntitynamePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(brand.BMBrand)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMEntitynamePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval brand.BMBrand = b.BrickInstance().Pr.(brand.BMBrand)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
