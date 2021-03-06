package attendee

import (
	"errors"
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

func (bd *BMAttendeeGuardianRS) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BMAttendeeGuardianRS) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "attendeeId"
	eq1.Vy = bd.AttendeeId
	req := request.Request{}
	req.Res = "BMAttendeeGuardianRS"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BMAttendeeGuardianRS
	err := bind.DeleteAll(c.(request.Request))
	return err
}

func (bd *BMAttendeeGuardianRS) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "attendeeId"
	eq1.Vy = bd.AttendeeId
	eq2 := request.Eqcond{}
	eq2.Ky = "guardianId"
	eq2.Vy = bd.GuardianId
	req := request.Request{}
	req.Res = "BMAttendeeGuardianRS"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BMAttendeeGuardianRS
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		return errors.New("Already Exist!")
	}
	return err
}
