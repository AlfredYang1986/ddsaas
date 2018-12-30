package reservablepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/reservable"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"time"
)

type BmReservablePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmReservablePushBrick) Exec() error {
	tmp := b.bk.Pr.(reservable.BmReservable)
	tmp.CreateTime = float64(time.Now().UnixNano() / 1e6)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmReservablePushBrick) Prepare(pr interface{}) error {
	req := pr.(reservable.BmReservable)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmReservablePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmReservablePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmReservablePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(reservable.BmReservable)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmReservablePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(reservable.BmReservable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

