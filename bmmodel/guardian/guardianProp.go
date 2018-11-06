package guardian

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMGuardianProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	GuardianId string `json:"guardianId" bson:"guardianId"`
	PersonId   string `json:"personId" bson:"personId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMGuardianProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMGuardianProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMGuardianProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMGuardianProp) QueryId() string {
	return bd.Id
}

func (bd *BMGuardianProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMGuardianProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMGuardianProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMGuardianProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMGuardianProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMGuardianProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMGuardianProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
