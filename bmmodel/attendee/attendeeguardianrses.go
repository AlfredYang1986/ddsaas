package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMAttendeeGuardianRSeS struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AgRsArr []BMAttendeeGuardianRS `json:"agrsarr" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRSeS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAttendeeGuardianRSeS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRSeS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAttendeeGuardianRSeS) QueryId() string {
	return bd.Id
}

func (bd *BMAttendeeGuardianRSeS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAttendeeGuardianRSeS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAttendeeGuardianRSeS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAttendeeGuardianRSeS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRSeS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAttendeeGuardianRSeS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAttendeeGuardianRSeS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMAttendeeGuardianRSeS) FindMulti(req request.Request) error {
	var agrses []BMAttendeeGuardianRS
	err := bmmodel.FindMutil(req, &agrses)
	bd.AgRsArr = agrses
	return err
}