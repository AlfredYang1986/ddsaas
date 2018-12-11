package applypush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/apply"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmApplyPushProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmApplyPushProp) Exec() error {
	tmp := b.bk.Pr.(apply.BmApply)

	one := tmp.Applyee
	oneBindOne := apply.BmApplyBindApplyee{}
	oneBindOne.ApplyId = tmp.Id
	oneBindOne.ApplyeeId = one.Id
	oneBindOne.Id_ = bson.NewObjectId()
	oneBindOne.Id = oneBindOne.Id_.Hex()
	oneBindOne.CheckExist()
	err := oneBindOne.InsertBMObject()

	sbt := apply.BmApplyBindKid{}
	sbt.Id_ = bson.NewObjectId()
	sbt.Id = sbt.Id_.Hex()
	sbt.ApplyId = tmp.Id
	sbt.Clear()
	for _, item := range tmp.Kids {
		itb := apply.BmApplyBindKid{}
		itb.Id_ = bson.NewObjectId()
		itb.Id = itb.Id_.Hex()
		itb.ApplyId = tmp.Id
		itb.KidId = item.Id
		//itb.CheckExist()
		itb.InsertBMObject()
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmApplyPushProp) Prepare(pr interface{}) error {
	req := pr.(apply.BmApply)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmApplyPushProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmApplyPushProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmApplyPushProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(apply.BmApply)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmApplyPushProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval apply.BmApply = b.BrickInstance().Pr.(apply.BmApply)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
