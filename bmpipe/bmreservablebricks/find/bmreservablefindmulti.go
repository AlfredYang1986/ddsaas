package reservablefind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/reservable"
	"io"
	"net/http"
)

type BmReservableFindMulti struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmReservableFindMulti) Exec() error {
	var tmp reservable.BmReservables
	err := tmp.FindMulti(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BmReservableFindMulti) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmReservableFindMulti) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmReservableFindMulti) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmReservableFindMulti) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(reservable.BmReservables)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmReservableFindMulti) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		results := b.BrickInstance().Pr.(reservable.BmReservables)
		jsonapi.ToJsonAPI(results.Reservables, w)
	}
}
