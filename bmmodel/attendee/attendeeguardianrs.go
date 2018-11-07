package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
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

func (bd *BMAttendeeGuardianRS) GetAttendee() (error, BmAttendee) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(bd.AttendeeId)
	req := request.Request{}
	req.Res = "BmAttendee"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	var attendee BmAttendee
	err := attendee.FindOne(c.(request.Request))
	return err, attendee
}

func (bd *BMAttendeeGuardianRS) GetGuardian() (error, guardian.BmGuardian) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(bd.GuardianId)
	req := request.Request{}
	req.Res = "BmGuardian"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	var guardian guardian.BmGuardian
	err := guardian.FindOne(c.(request.Request))
	guardian.ReSetPerson()
	return err, guardian
}
