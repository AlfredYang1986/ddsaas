package contact

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

/*
   Replace entityname && Entityname
   Define Attibute1/2/... && attibute1/2/...
   Case-sensitive
*/

type BMContactProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ContactId string   `json:"contact_id" bson:"contact_id"`
	OrderIds  []string `json:"order_ids" bson:"order_ids"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMContactProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMContactProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMContactProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMContactProp) QueryId() string {
	return bd.Id
}

func (bd *BMContactProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMContactProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMContactProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMContactProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMContactProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMContactProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMContactProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
