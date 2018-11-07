package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMTeacherProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	TeacherId string `json:"teacherId" bson:"teacherId"`
	PersonId   string `json:"personId" bson:"personId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMTeacherProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMTeacherProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMTeacherProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMTeacherProp) QueryId() string {
	return bd.Id
}

func (bd *BMTeacherProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMTeacherProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMTeacherProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMTeacherProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMTeacherProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMTeacherProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMTeacherProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMTeacherProp) GetTeacher() (error, BmTeacher) {
	eq := request.EqCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(bd.TeacherId)
	req := request.Request{}
	req.Res = "BmTeacher"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	var tech BmTeacher
	err := tech.FindOne(c.(request.Request))
	return err, tech
}

func (bd *BMTeacherProp) GetPerson() (error, person.BmPerson) {
	eq := request.EqCond{}
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