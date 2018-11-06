package applyee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMApplyee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	RelationShip string                `json:"relation_ship" bson:"relation_ship"`
	Attendees    []attendee.BMAttendee `json:"attendees" bson:"relationships"`
	Person       person.BMPerson       `json:"person" bson:"relationships"`
}

/*-----------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMApplyee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMApplyee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMApplyee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMApplyee) QueryId() string {
	return bd.Id
}

func (bd *BMApplyee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMApplyee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMApplyee) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "person":
		bd.Person = v.(person.BMPerson)
	case "attendee":
		var rst []attendee.BMAttendee
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(attendee.BMAttendee))
		}
		bd.Attendees = rst
	}

	return bd
}

func (bd BMApplyee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMApplyee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMApplyee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMApplyee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
