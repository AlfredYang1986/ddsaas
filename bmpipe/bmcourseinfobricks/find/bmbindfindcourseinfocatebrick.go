package courseinfofind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmBindFindSessionInfoCatBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBindFindSessionInfoCatBrick) Exec() error {
	var tmp sessioninfo.BmSessionInfo = b.BrickInstance().Pr.(sessioninfo.BmSessionInfo)

	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BmSessionBindCat"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := sessioninfo.BmSessionBindCat{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.CategoryId
	req0 := request.Request{}
	req0.Res = "BmCategory"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := category.BmCategory{}
	err = result.FindOne(c0.(request.Request))
	tmp.Cate = result

	imgs, err := b.findImgs()
	tmp.TagImgs = imgs

	b.BrickInstance().Pr = tmp
	return err
}

func (b *BmBindFindSessionInfoCatBrick) Prepare(pr interface{}) error {
	req := pr.(sessioninfo.BmSessionInfo)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBindFindSessionInfoCatBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBindFindSessionInfoCatBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBindFindSessionInfoCatBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessioninfo.BmSessionInfo)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBindFindSessionInfoCatBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(sessioninfo.BmSessionInfo)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func (b *BmBindFindSessionInfoCatBrick) findImgs() ([]tagimg.BmTagImg, error) {
	si := b.BrickInstance().Pr.(sessioninfo.BmSessionInfo)

	req := request.Request{}
	req.Res = "BmBindSessionImg"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = si.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []sessioninfo.BmBindSessionImg
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return nil, err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.TagImgId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var imgs []tagimg.BmTagImg
	err = bmmodel.FindMutilWithBson("BmTagImg", or_condi, &imgs)

	for i, ir := range imgs {
		ir.ResetIdWithId_()
		imgs[i] = ir
	}

	return imgs, err
}
