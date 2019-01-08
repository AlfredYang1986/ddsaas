package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindBrandCategory struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	BrandId    string `json:"brandId" bson:"brandId"`
	CategoryId string `json:"categoryId" bson:"categoryId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindBrandCategory) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindBrandCategory) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindBrandCategory) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindBrandCategory) QueryId() string {
	return bd.Id
}

func (bd *BmBindBrandCategory) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindBrandCategory) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindBrandCategory) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindBrandCategory) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindBrandCategory) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindBrandCategory) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindBrandCategory) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindBrandCategory) CheckExist() error {

	eq2 := request.Eqcond{}
	eq2.Ky = "brandId"
	eq2.Vy = bd.BrandId
	req := request.Request{}
	req.Res = "BmBindBrandCategory"
	var condi []interface{}
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindBrandCategory
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}