package yardpush

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"gopkg.in/mgo.v2/bson"
)

type BmBindYardPropBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBindYardPropBrick) Exec() error {
	tmp := b.bk.Pr.(yard.BmYard)

	for _, item := range tmp.Rooms {
		ist := yard.BmBindYardRoom{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.RoomId = item.Id
		ist.YardId = tmp.Id
		ist.CheckExist()
		ist.InsertBMObject()
	}

	for _, item := range tmp.TagImgs {
		ist := yard.BmBindYardImg{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.YardId = tmp.Id
		ist.TagImgId = item.Id
		ist.CheckExist()
		ist.InsertBMObject()
	}

	for _, item := range tmp.Certifications {
		ist := yard.BmBindYardCertific{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.YardId = tmp.Id
		ist.CertificationId = item.Id
		ist.InsertBMObject()
	}

	b.bk.Pr = tmp
	return nil
}

func (b *BmBindYardPropBrick) Prepare(pr interface{}) error {
	req := pr.(yard.BmYard)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBindYardPropBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBindYardPropBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBindYardPropBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(yard.BmYard)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBindYardPropBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(yard.BmYard)
		jsonapi.ToJsonAPI(&reval, w)
	}
}