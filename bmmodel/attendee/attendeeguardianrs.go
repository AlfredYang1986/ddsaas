package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMAttendeeGuardianRS struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AttendeeId string `json:"attendeeId" bson:"attendeeId"`
	GuardianId string `json:"guardianId" bson:"guardianId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAttendeeGuardianRS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAttendeeGuardianRS) QueryId() string {
	return bd.Id
}

func (bd *BMAttendeeGuardianRS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAttendeeGuardianRS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAttendeeGuardianRS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAttendeeGuardianRS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAttendeeGuardianRS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAttendeeGuardianRS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAttendeeGuardianRS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
