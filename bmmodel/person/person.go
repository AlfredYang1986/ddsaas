package person

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmPerson struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name     string `json:"name" bson:"name"`
	Nickname string `json:"nickname" bson:"nickname"`
	Icon     string `json:"icon" bson:"icon"`
	Dob      int64 `json:"dob" bson:"dob"`
	Gender   int64 `json:"gender" bson:"gender"`
	RegDate  int64 `json:"reg_date" bson:"reg_date"`

	//Address address.BmAddress `json:"Address" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmPerson) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmPerson) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmPerson) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmPerson) QueryId() string {
	return bd.Id
}

func (bd *BmPerson) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmPerson) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmPerson) SetConnect(tag string, v interface{}) interface{} {
	//switch tag {
	//case "address":
	//	bd.Address = v.(address.BmAddress)
	//}
	return bd
}

func (bd BmPerson) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmPerson) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmPerson) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmPerson) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
