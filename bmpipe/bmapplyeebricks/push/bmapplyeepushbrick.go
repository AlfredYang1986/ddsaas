package applyeepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"io"
	"net/http"
)

type BmApplyeePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyeePushBrick) Exec() error {
	var err error
	tmp := b.bk.Pr.(applyee.BmApplyee)

	if tmp.Id != "" && tmp.Id_.Valid() && tmp.Valid() {
		//TODO:现在[push已存在的]相当于[update]
		tmp.CheckExist()
		err = tmp.InsertBMObject()
		b.bk.Pr = tmp
		//if  tmp.IsRegisted() {
		//	b.bk.Err = -9
		//}
	} else {
		b.BrickInstance().Err = -10
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmApplyeePushBrick) Prepare(pr interface{}) error {
	req := pr.(applyee.BmApplyee)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyeePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	ec := b.BrickInstance().Err
	if int(idx) < tmp-1 && ec == 0 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyeePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyeePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(applyee.BmApplyee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyeePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(auth.BmLoginSucceed)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

