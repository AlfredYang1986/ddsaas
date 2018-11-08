package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmTeachers struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Teachers []BmTeacher `json:"teachers" jsonapi:"relationships"`
}

type BmTeachersResult struct {
	Id          string   `json:"id"`
	TeacherIds []string `json:"teacherids"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmTeachers) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmTeachers) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmTeachers) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmTeachers) QueryId() string {
	return bd.Id
}

func (bd *BmTeachers) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmTeachers) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmTeachers) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "teachers":
		var rst []BmTeacher
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BmTeacher))
		}
		bd.Teachers = rst
	}
	return bd
}

func (bd BmTeachers) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmTeachers) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmTeachers) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmTeachers) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmTeachers) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Teachers)
	for i, r := range bd.Teachers {
		r.ResetIdWithId_()
		bd.Teachers[i] = r
	}
	return err
}

//func (bd *BmTeachers) ReSetPerson() error {
//
//	var err error
//	for i, r := range bd.Teachers {
//		attendeeId := r.Id
//		eq := request.Eqcond{}
//		eq.Ky = "attendeeId"
//		eq.Vy = attendeeId
//		req1 := request.Request{}
//		req1.Res = "BMAttendeeProp"
//		var condi1 []interface{}
//		condi1 = append(condi1, eq)
//		c1 := req1.SetConnect("conditions", condi1)
//		var attendeeProp BMTeacherProp
//		err = attendeeProp.FindOne(c1.(request.Request))
//		if err != nil {
//			return err
//		}
//
//		err, person := attendeeProp.GetPerson()
//		if err != nil {
//			return err
//		}
//		r.Person = person
//
//		bd.Teachers[i] = r
//	}
//	return err
//}