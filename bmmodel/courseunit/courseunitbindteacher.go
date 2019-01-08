package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnitBindTeacher struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseUnitId string `json:"courseUnitId" bson:"courseUnitId"`
	TeacherId    string `json:"teacherId" bson:"teacherId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindTeacher) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnitBindTeacher) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindTeacher) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnitBindTeacher) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnitBindTeacher) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnitBindTeacher) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnitBindTeacher) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCourseUnitBindTeacher) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindTeacher) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnitBindTeacher) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCourseUnitBindTeacher) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnitBindTeacher) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnitBindTeacher) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmCourseUnitBindTeacher) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "courseUnitId"
	eq1.Vy = bd.CourseUnitId
	req := request.Request{}
	req.Res = "BmCourseUnitBindTeacher"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmCourseUnitBindTeacher
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
