package sessioninfo

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionInfo struct {
	Id       string        `json:"id"`
	Id_      bson.ObjectId `bson:"_id"`
	BrandId  string        `json:"brandId" bson:"brandId"`
	Title    string        `json:"title" bson:"title"`
	Subtitle string        `json:"subtitle" bson:"subtitle"`
	Alb      float64       `json:"alb" bson:"alb"`
	Aub      float64       `json:"aub" bson:"aub"`
	Level    string        `json:"level" bson:"level"`
	Count    float64       `json:"count" bson:"count"`
	Length   float64       `json:"length" bson:"length"`

	Cate category.BmCategory `json:"Cate" jsonapi:"relationships"`

	//TODO:20181109新增的
	Description  string            `json:"description" bson:"description"`
	Harvest      string            `json:"harvest" bson:"harvest"`
	Acquisition  string            `json:"acquisition" bson:"acquisition"`
	Accompany    float64           `json:"accompany" bson:"accompany"`
	Status       float64           `json:"status" bson:"status"` //0活动 1体验课 2普通课程
	Including    string            `json:"inc" bson:"including"`
	Carrying     string            `json:"carrying" bson:"carrying"`
	Notice       string            `json:"notice" bson:"notice"`
	PlayChildren string            `json:"play_children" bson:"play_children"`
	Cover        string            `json:"cover" bson:"cover"`
	TagImgs      []tagimg.BmTagImg `json:"Tagimgs" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionInfo) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionInfo) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionInfo) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionInfo) QueryId() string {
	return bd.Id
}

func (bd *BmSessionInfo) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionInfo) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionInfo) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Cate":
		bd.Cate = v.(category.BmCategory)
	case "Tagimgs":
		var rst []tagimg.BmTagImg
		for _, item := range v.([]interface{}) {
			tmp := item.(tagimg.BmTagImg)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.TagImgs = rst
	}
	return bd
}

func (bd BmSessionInfo) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionInfo) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionInfo) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionInfo) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionInfo) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmSessionInfo) ReSetProp() error {
	bd.reSetCategory()
	bd.reSetTagImg()
	return nil
}

func (bd *BmSessionInfo) DeleteProp() error {
    bd.deleteBmSessionBindCat()
    bd.deleteBmBindSessionImg()
	return nil
}

func (bd *BmSessionInfo) reSetCategory() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionBindCat"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionBindCat{}
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
	bd.Cate = result

	return err
}

func (bd *BmSessionInfo) reSetTagImg() error {

	req := request.Request{}
	req.Res = "BmBindSessionImg"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindSessionImg
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
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
	bd.TagImgs = imgs

	return err
}

func (bd *BmSessionInfo) deleteBmSessionBindCat() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionBindCat"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionBindCat{}
	err := reval.FindOne(c.(request.Request))
	err = reval.DeleteAll(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.CategoryId

	if eq0.Vy == "" {
		return err
	}

	req0 := request.Request{}
	req0.Res = "BmCategory"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := category.BmCategory{}
	err = result.DeleteAll(c0.(request.Request))
	bd.Cate = result

	return err
}

func (bd *BmSessionInfo) deleteBmBindSessionImg() error {

	req := request.Request{}
	req.Res = "BmBindSessionImg"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var bind BmBindSessionImg
	err := bind.DeleteAll(c.(request.Request))

	var reval []BmBindSessionImg
	err = bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.TagImgId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	err = bmmodel.DeleteMutilWithBson("BmTagImg", or_condi)

	return err
}
