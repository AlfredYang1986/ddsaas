package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMAttendeeProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AttendeeId string `json:"attendeeId" bson:"attendeeId"`
	PersonId   string `json:"personId" bson:"personId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAttendeeProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAttendeeProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAttendeeProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAttendeeProp) QueryId() string {
	return bd.Id
}

func (bd *BMAttendeeProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAttendeeProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAttendeeProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAttendeeProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAttendeeProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAttendeeProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAttendeeProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMAttendeeProp) GetAttendee() (error, BmAttendee) {
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

func (bd *BMAttendeeProp) GetPerson() (error, person.BmPerson) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(bd.PersonId)
	req := request.Request{}
	req.Res = "BmPerson"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	person := person.BmPerson{}
	err := person.FindOne(c.(request.Request))
	return err, person
}
