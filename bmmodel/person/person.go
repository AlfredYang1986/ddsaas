package person

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/address"
	"gopkg.in/mgo.v2/bson"
)

type BMPerson struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name     string `json:"name" bson:"name"`
	Nickname string `json:"nickname" bson:"nickname"`
	Icon     string `json:"icon" bson:"icon"`
	Dob      int64 `json:"dob" bson:"dob"`
	Gender   string `json:"gender" bson:"gender"`
	RegDate  int64 `json:"reg_date" bson:"reg_date"`

	Address address.BMAddress `json:"address" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMPerson) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMPerson) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMPerson) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMPerson) QueryId() string {
	return bd.Id
}

func (bd *BMPerson) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMPerson) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMPerson) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "address":
		bd.Address = v.(address.BMAddress)
	}
	return bd
}

func (bd BMPerson) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMPerson) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMPerson) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMPerson) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
