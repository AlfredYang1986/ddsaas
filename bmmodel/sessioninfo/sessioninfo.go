package sessioninfo

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionInfo struct {
	Id      string        `json:"id"`
	Id_     bson.ObjectId `bson:"_id"`
	BrandId string        `json:"brandId" bson:"brandId"`
	Title   string        `json:"title" bson:"title"`
	Alb     float64       `json:"alb" bson:"alb"`
	Aub     float64       `json:"aub" bson:"aub"`
	Level   string        `json:"level" bson:"level"`
	Count   float64       `json:"count" bson:"count"`
	Length  float64       `json:"length" bson:"length"`

	Cate category.BmCategory `json:"Cate" jsonapi:"relationships"`

	//TODO:20181109新增的
	Description string `json:"description" bson:"description"`
	Harvest     string `json:"harvest" bson:"harvest"`
	Acquisition string `json:"acquisition" bson:"acquisition"`
	Accompany   int64  `json:"accompany" bson:"accompany"`
	Including   string `json:"including" bson:"including"`
	Carrying    string `json:"carrying" bson:"carrying"`
	Notice      string `json:"notice" bson:"notice"`
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
