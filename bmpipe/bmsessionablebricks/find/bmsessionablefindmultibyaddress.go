package sessionablefind

import (
	"errors"
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

type BmSessionableFindMultiByYard struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmSessionableFindMultiByYard) Exec() error {
	var tmp sessionable.BmSessionables
	var result []sessionable.BmSessionable
	yardId := ""
	for _, cond := range b.bk.Req.Eqcond {
		if cond.Ky == "id" && cond.Ct == "BmYard" {
			yardId = cond.Vy.(string)
		}
	}
	if yardId == "" {
		return errors.New("no yard id")
	}
	err := tmp.FindMulti(*b.bk.Req)
	for _, v := range tmp.Sessionables {
		if v.Yard.Id == yardId {
			result = append(result, v)
		}
	}
	tmp.Sessionables = result
	b.bk.Pr = tmp
	return err
}

func (b *BmSessionableFindMultiByYard) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmSessionableFindMultiByYard) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmSessionableFindMultiByYard) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmSessionableFindMultiByYard) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessionable.BmSessionables)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmSessionableFindMultiByYard) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		results := b.BrickInstance().Pr.(sessionable.BmSessionables)
		jsonapi.ToJsonAPI(results.Sessionables, w)
	}
}
