package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmAttendees struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Attendees []BmAttendee `json:"attendees" jsonapi:"relationships"`
}

type BmAttendeesResult struct {
	Id          string   `json:"id"`
	AttendeeIds []string `json:"attendeeids"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmAttendees) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmAttendees) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmAttendees) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmAttendees) QueryId() string {
	return bd.Id
}

func (bd *BmAttendees) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmAttendees) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmAttendees) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmAttendees) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmAttendees) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmAttendees) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAttendees) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmAttendees) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Attendees)
	for i, r := range bd.Attendees {
		r.ResetIdWithId_()
		bd.Attendees[i] = r
	}
	return err
}

func (bd *BmAttendees) ReSetPerson() error {

	var err error
	for i, r := range bd.Attendees {
		attendeeId := r.Id
		eq := request.EqCond{}
		eq.Ky = "attendeeId"
		eq.Vy = attendeeId
		req1 := request.Request{}
		req1.Res = "BMAttendeeProp"
		var condi1 []interface{}
		condi1 = append(condi1, eq)
		c1 := req1.SetConnect("conditions", condi1)
		var attendeeProp BMAttendeeProp
		err = attendeeProp.FindOne(c1.(request.Request))
		if err != nil {
			return err
		}

		err, person := attendeeProp.GetPerson()
		if err != nil {
			return err
		}
		r.Person = person

		bd.Attendees[i] = r
	}
	return err
}

