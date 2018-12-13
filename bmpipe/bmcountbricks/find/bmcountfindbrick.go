package countfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/count"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/hashicorp/go-uuid"
	"net/http"
	"io"
)

type BmCountFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmCountFindBrick) Exec() error {
	var tmp count.BmCount
	n, err := tmp.FindCount(*b.bk.Req)
	uui, _ := uuid.GenerateUUID()
	tmp.Id = uui
	tmp.Res = b.bk.Req.Res
	tmp.Count = n
	b.bk.Pr = tmp
	return err
}

func (b *BmCountFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmCountFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmCountFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmCountFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(count.BmCount)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmCountFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval count.BmCount = b.BrickInstance().Pr.(count.BmCount)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
