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

type BmApplyPushKids struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyPushKids) Exec() error {
	tmp := b.bk.Pr.(apply.BmApply)
	for _, k := range tmp.Kids {
		k.InsertBMObject()
	}
	b.bk.Pr = tmp
	return nil
}

func (b *BmApplyPushKids) Prepare(pr interface{}) error {
	req := pr.(apply.BmApply)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyPushKids) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyPushKids) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyPushKids) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(apply.BmApply)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyPushKids) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval apply.BmApply = b.BrickInstance().Pr.(apply.BmApply)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

