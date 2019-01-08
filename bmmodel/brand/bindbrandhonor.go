package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindBrandHonor struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	BrandId string `json:"brandId" bson:"brandId"`
	HonorId string `json:"honorId" bson:"honorId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindBrandHonor) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindBrandHonor) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindBrandHonor) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindBrandHonor) QueryId() string {
	return bd.Id
}

func (bd *BmBindBrandHonor) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindBrandHonor) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindBrandHonor) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindBrandHonor) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindBrandHonor) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindBrandHonor) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindBrandHonor) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindBrandHonor) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmBindBrandHonor) Clear() error {

	eq2 := request.Eqcond{}
	eq2.Ky = "brandId"
	eq2.Vy = bd.BrandId
	req := request.Request{}
	req.Res = "BmBindBrandHonor"
	var condi []interface{}
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindBrandHonor
	err := bind.DeleteAll(c.(request.Request))
	return err
}
