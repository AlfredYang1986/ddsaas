package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnitBindSessionable struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseUnitId  string `json:"courseUnitId" bson:"courseUnitId"`
	SessionableId string `json:"sessionableId" bson:"sessionableId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindSessionable) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnitBindSessionable) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindSessionable) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnitBindSessionable) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnitBindSessionable) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnitBindSessionable) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnitBindSessionable) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCourseUnitBindSessionable) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindSessionable) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnitBindSessionable) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCourseUnitBindSessionable) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnitBindSessionable) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnitBindSessionable) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmCourseUnitBindSessionable) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "courseUnitId"
	eq1.Vy = bd.CourseUnitId
	req := request.Request{}
	req.Res = "BmCourseUnitBindSessionable"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmCourseUnitBindSessionable
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
