package kid

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmKid struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name         string  `json:"name" bson:"name"`
	NickName     string  `json:"nickname" bson:"nickname"`
	Gender       float64 `json:"gender" bson:"gender"`
	Dob          float64 `json:"dob" bson:"dob"`
	GuardianRole float64 `json:"guardian_role" bson:"guardian_role"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmKid) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmKid) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmKid) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmKid) QueryId() string {
	return bd.Id
}

func (bd *BmKid) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmKid) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmKid) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmKid) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmKid) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmKid) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmKid) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmKid) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
