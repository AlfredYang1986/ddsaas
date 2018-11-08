package yardfind

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
)

type BmYardFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmYardFindBrick) Exec() error {
	var tmp yard.BmYard
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BmYardFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmYardFindBrick) Done(pkg string, idx int64, e error) error {

	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmYardFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmYardFindBrick) ResultTo(w io.Writer) error {

	pr := b.BrickInstance().Pr
	tmp := pr.(yard.BmYard)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmYardFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		pr := b.BrickInstance().Pr
		tmp := pr.(yard.BmYard)
		jsonapi.ToJsonAPI(&tmp, w)
	}
}
