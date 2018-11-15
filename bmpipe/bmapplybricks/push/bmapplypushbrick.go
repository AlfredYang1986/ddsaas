package applypush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/apply"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BmApplyPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyPushBrick) Exec() error {
	tmp := b.bk.Pr.(apply.BmApply)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmApplyPushBrick) Prepare(pr interface{}) error {
	req := pr.(apply.BmApply)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(apply.BmApply)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval apply.BmApply = b.BrickInstance().Pr.(apply.BmApply)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

