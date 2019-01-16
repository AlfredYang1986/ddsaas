package sessionablefind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"io"
	"net/http"
)

type BmSessionableFindMulti struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmSessionableFindMulti) Exec() error {
	var tmp sessionable.BmSessionables
	err := tmp.FindMulti(*b.bk.Req)
	for i, v := range tmp.Sessionables[:] {
		ReSetClassDate(&v)
		tmp.Sessionables[i] = v
	}
	b.bk.Pr = tmp
	return err
}

func (b *BmSessionableFindMulti) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmSessionableFindMulti) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmSessionableFindMulti) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmSessionableFindMulti) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessionable.BmSessionables)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmSessionableFindMulti) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		results := b.BrickInstance().Pr.(sessionable.BmSessionables)
		jsonapi.ToJsonAPI(results.Sessionables, w)
	}
}
