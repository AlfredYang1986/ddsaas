package sessionablepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"time"
)

type BmSessionablePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmSessionablePushBrick) Exec() error {
	tmp := b.bk.Pr.(sessionable.BmSessionable)
	tmp.CreateTime = float64(time.Now().UnixNano() / 1e6)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmSessionablePushBrick) Prepare(pr interface{}) error {
	req := pr.(sessionable.BmSessionable)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmSessionablePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmSessionablePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmSessionablePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessionable.BmSessionable)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmSessionablePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval sessionable.BmSessionable = b.BrickInstance().Pr.(sessionable.BmSessionable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

