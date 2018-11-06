package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/payment"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMAttendee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Intro       string `json:"intro" bson:"intro"`
	Status      string `json:"status" bson:"status"`
	LessonCount int64 `json:"lesson_count" bson:"lesson_count"`

	Person    person.BMPerson       `json:"person" jsonapi:"relationships"`
	Guardians []guardian.BMGuardian `json:"guardians" jsonapi:"relationships"`
	Payments  []payment.BMPayment   `json:"payments" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAttendee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAttendee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAttendee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAttendee) QueryId() string {
	return bd.Id
}

func (bd *BMAttendee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAttendee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAttendee) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "person":
		bd.Person = v.(person.BMPerson)
	case "guardians":
		var rst []guardian.BMGuardian
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(guardian.BMGuardian))
		}
		bd.Guardians = rst
	case "payments":
		var rst []payment.BMPayment
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(payment.BMPayment))
		}
		bd.Payments = rst
	}
	return bd
}

func (bd BMAttendee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAttendee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAttendee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAttendee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
