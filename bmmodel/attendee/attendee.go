package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/payment"
	"gopkg.in/mgo.v2/bson"
)

type BmAttendee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Intro       string `json:"intro" bson:"intro"`
	Status      string `json:"status" bson:"status"`
	LessonCount int64 `json:"lesson_count" bson:"lesson_count"`

	Name     string `json:"name" bson:"name"`
	Nickname string `json:"nickname" bson:"nickname"`
	Icon     string `json:"icon" bson:"icon"`
	Dob      int64 `json:"dob" bson:"dob"`
	Gender   int64 `json:"gender" bson:"gender"`
	RegDate  int64 `json:"reg_date" bson:"reg_date"`
	Contact  string `json:"contact" bson:"contact"`
	WeChat  string `json:"wechat" bson:"wechat"`

	//Person    person.BmPerson       `json:"Person" jsonapi:"relationships"`
	Guardians []guardian.BmGuardian `json:"Guardians" jsonapi:"relationships"`
	Payments  []payment.BMPayment   `json:"Payments" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmAttendee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmAttendee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmAttendee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmAttendee) QueryId() string {
	return bd.Id
}

func (bd *BmAttendee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmAttendee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmAttendee) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	//case "Person":
	//	bd.Person = v.(person.BmPerson)
	case "Guardians":
		var rst []guardian.BmGuardian
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(guardian.BmGuardian))
		}
		bd.Guardians = rst
	case "Payments":
		var rst []payment.BMPayment
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(payment.BMPayment))
		}
		bd.Payments = rst
	}
	return bd
}

func (bd BmAttendee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmAttendee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmAttendee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAttendee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmAttendee) GetAttendeeProp() (error, BMAttendeeProp) {

	eq := request.Eqcond{}
	eq.Ky = "attendeeId"
	eq.Vy = bd.Id
	req1 := request.Request{}
	req1.Res = "BMAttendeeProp"
	var condi1 []interface{}
	condi1 = append(condi1, eq)
	c1 := req1.SetConnect("conditions", condi1)
	attendeeProp := BMAttendeeProp{}
	err := attendeeProp.FindOne(c1.(request.Request))
	return err, attendeeProp

}