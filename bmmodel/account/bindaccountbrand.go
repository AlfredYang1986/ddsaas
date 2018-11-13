package account

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindAccountBrand struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AccountId string `json:"accountId" bson:"accountId"`
	BrandId   string `json:"brandId" bson:"brandId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindAccountBrand) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindAccountBrand) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindAccountBrand) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindAccountBrand) QueryId() string {
	return bd.Id
}

func (bd *BmBindAccountBrand) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindAccountBrand) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindAccountBrand) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindAccountBrand) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindAccountBrand) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindAccountBrand) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindAccountBrand) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindAccountBrand) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "accountId"
	eq1.Vy = bd.AccountId
	eq2 := request.Eqcond{}
	eq2.Ky = "brandId"
	eq2.Vy = bd.BrandId
	req := request.Request{}
	req.Res = "BmBindAccountBrand"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindAccountBrand
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
