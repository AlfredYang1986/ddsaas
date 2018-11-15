package applyeefind

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"io"
	"net/http"
)

type BmApplyeeFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyeeFindBrick) Exec() error {
	var tmp applyee.BmApplyee
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BmApplyeeFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmApplyeeFindBrick) Done(pkg string, idx int64, e error) error {
	if e != nil && e.Error() == "not found" {
		reval := applyee.BmApplyee{}
		reval.WeChatOpenid = b.BrickInstance().Req.CondiQueryVal("wechat_openid", "BmApplyee").(string)
		b.BrickInstance().Pr = reval
		bmrouter.NextBrickRemote("pushapplyee", 0, b)
	} else {
		bmrouter.NextBrickRemote("applyeegeneratetoken", 0, b)
	}
	return nil
}

func (b *BmApplyeeFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyeeFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(applyee.BmApplyee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyeeFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(auth.BmLoginSucceed)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

