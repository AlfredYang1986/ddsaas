package guardian

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMGuardian struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Person person.BMPerson 	`json:"person" jsonapi:"relationships"`
	Relationship string `json:"relation_ship" bson:"relation_ship"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMGuardian) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMGuardian) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMGuardian) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMGuardian) QueryId() string {
	return bd.Id
}

func (bd *BMGuardian) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMGuardian) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMGuardian) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "person":
		bd.Person = v.(person.BMPerson)
	}
	return bd
}

func (bd BMGuardian) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMGuardian) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMGuardian) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMGuardian) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
