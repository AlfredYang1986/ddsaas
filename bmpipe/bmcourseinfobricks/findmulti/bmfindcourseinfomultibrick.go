package courseinfofindmulti

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
)

type BmFindSessionInfoMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmFindSessionInfoMultiBrick) Exec() error {
	var tmp sessioninfo.BmSessionInfos
	err := tmp.FindMulti(b.BrickInstance().Pr.(request.Request))
	b.bk.Pr = tmp
	return err
}

func (b *BmFindSessionInfoMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmFindSessionInfoMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmFindSessionInfoMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmFindSessionInfoMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessioninfo.BmSessionInfos)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmFindSessionInfoMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(sessioninfo.BmSessionInfos)
		jsonapi.ToJsonAPI(reval.Sessions, w)
	}
}