package yardfind

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
)

type BmYardFindBindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmYardFindBindBrick) Exec() error {
	tmp := b.bk.Pr.(yard.BmYard)
	//err := tmp.FindOne(*b.bk.Req)

	rooms, err := b.findRooms()
	if err == nil {
		tmp.Rooms = rooms
	}

	imgs, err := b.findImgs()
	if err == nil {
		tmp.TagImgs = imgs
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmYardFindBindBrick) Prepare(pr interface{}) error {
	req := pr.(yard.BmYard)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmYardFindBindBrick) Done(pkg string, idx int64, e error) error {

	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmYardFindBindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmYardFindBindBrick) ResultTo(w io.Writer) error {

	pr := b.BrickInstance().Pr
	tmp := pr.(yard.BmYard)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmYardFindBindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		pr := b.BrickInstance().Pr
		tmp := pr.(yard.BmYard)
		jsonapi.ToJsonAPI(&tmp, w)
	}
}

func (b BmYardFindBindBrick) findImgs() ([]tagimg.BmTagImg, error) {
	yd := b.BrickInstance().Pr.(yard.BmYard)

	req := request.Request{}
	req.Res = "BmBindYardImg"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "yardId"
	eq.Vy = yd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []yard.BmBindYardImg
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

func (b BmYardFindBindBrick) findRooms() ([]room.BmRoom, error){
	yd := b.BrickInstance().Pr.(yard.BmYard)

	req := request.Request{}
	req.Res = "BmBindYardRoom"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "yardId"
	eq.Vy = yd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []yard.BmBindYardRoom
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return nil, err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.RoomId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var rooms []room.BmRoom
	err = bmmodel.FindMutilWithBson("BmRoom", or_condi, &rooms)

	for i, ir := range rooms {
		ir.ResetIdWithId_()
		rooms[i] = ir
	}

	return rooms, err
}
