package tagimgpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BmTagImgPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTagImgPushBrick) Exec() error {
	tmp := b.bk.Pr.(tagimg.BmTagImg)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmTagImgPushBrick) Prepare(pr interface{}) error {
	req := pr.(tagimg.BmTagImg)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTagImgPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTagImgPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTagImgPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(tagimg.BmTagImg)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTagImgPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval tagimg.BmTagImg = b.BrickInstance().Pr.(tagimg.BmTagImg)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

