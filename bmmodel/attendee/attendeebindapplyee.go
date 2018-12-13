package attendee

import (
	"errors"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMAttendeeBindApplyee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AttendeeId string `json:"attendeeId" bson:"attendeeId"`
	ApplyeeId string `json:"applyeeId" bson:"applyeeId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAttendeeBindApplyee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAttendeeBindApplyee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAttendeeBindApplyee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAttendeeBindApplyee) QueryId() string {
	return bd.Id
}

func (bd *BMAttendeeBindApplyee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAttendeeBindApplyee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAttendeeBindApplyee) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAttendeeBindApplyee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAttendeeBindApplyee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAttendeeBindApplyee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAttendeeBindApplyee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMAttendeeBindApplyee) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BMAttendeeBindApplyee) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "attendeeId"
	eq1.Vy = bd.AttendeeId
	req := request.Request{}
	req.Res = "BMAttendeeBindApplyee"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BMAttendeeBindApplyee
	err := bind.DeleteAll(c.(request.Request))
	return err
}

func (bd *BMAttendeeBindApplyee) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "attendeeId"
	eq1.Vy = bd.AttendeeId
	eq2 := request.Eqcond{}
	eq2.Ky = "applyeeId"
	eq2.Vy = bd.ApplyeeId
	req := request.Request{}
	req.Res = "BMAttendeeBindApplyee"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BMAttendeeBindApplyee
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		return errors.New("Already Exist!")
	}
	return err
}
