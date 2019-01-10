package applyfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/apply"
	"io"
	"net/http"
)

type BmApplyDetailFindMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyDetailFindMultiBrick) Exec() error {
	var tmp apply.BmApplyDetails
	err := tmp.FindMulti(b.BrickInstance().Pr.(request.Request))
	b.bk.Pr = tmp
	return err
}

func (b *BmApplyDetailFindMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyDetailFindMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyDetailFindMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyDetailFindMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(apply.BmApplyDetails)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyDetailFindMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(apply.BmApplyDetails)
		jsonapi.ToJsonAPI(reval.Applies, w)
	}
}