package applyfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/apply"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
)

type BmApplyDetailFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyDetailFindBrick) Exec() error {
	var tmp apply.BmApplyDetail
	err := tmp.FindOne(*b.bk.Req)
	tmp.ReSetProp()
	b.bk.Pr = tmp
	return err
}

func (b *BmApplyDetailFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmApplyDetailFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyDetailFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyDetailFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(apply.BmApplyDetail)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyDetailFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval apply.BmApplyDetail = b.BrickInstance().Pr.(apply.BmApplyDetail)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
