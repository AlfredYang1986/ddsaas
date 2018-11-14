package reservablepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/reservable"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"io"
	"net/http"
)

type BmReservablePushSession struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmReservablePushSession) Exec() error {
	tmp := b.bk.Pr.(reservable.BmReservable)
	session := tmp.SessionInfo
	err := session.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmReservablePushSession) Prepare(pr interface{}) error {
	req := pr.(reservable.BmReservable)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmReservablePushSession) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		pr := b.bk.Pr.(reservable.BmReservable)
		b.BrickInstance().Pr = pr.SessionInfo
		bmrouter.NextBrickRemote(pkg, idx+1, b)
		b.BrickInstance().Pr = pr
	}
	return nil
}

func (b *BmReservablePushSession) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmReservablePushSession) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessioninfo.BmSessionInfo)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmReservablePushSession) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(reservable.BmReservable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

