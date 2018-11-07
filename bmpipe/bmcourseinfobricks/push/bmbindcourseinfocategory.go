package courseinfopush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"gopkg.in/mgo.v2/bson"
)

type BmBindSessionCatPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBindSessionCatPushBrick) Exec() error {
	tmp := b.bk.Pr.(sessioninfo.BmSessionInfo)
	cat := tmp.Cat

	sbc := sessioninfo.BmSessionBindCat{}
	sbc.CategoryId = cat.Id
	sbc.SessionId = tmp.Id
	sbc.Id_ = bson.NewObjectId()
	sbc.Id = sbc.Id_.Hex()

	err := sbc.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmBindSessionCatPushBrick) Prepare(pr interface{}) error {
	req := pr.(sessioninfo.BmSessionInfo)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBindSessionCatPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBindSessionCatPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBindSessionCatPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessioninfo.BmSessionInfo)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBindSessionCatPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(sessioninfo.BmSessionInfo)
		jsonapi.ToJsonAPI(&reval, w)
	}
}