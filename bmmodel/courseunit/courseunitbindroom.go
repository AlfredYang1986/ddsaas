package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnitBindRoom struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseUnitId string `json:"courseUnitId" bson:"courseUnitId"`
	RoomId    string `json:"roomId" bson:"roomId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindRoom) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnitBindRoom) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindRoom) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnitBindRoom) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnitBindRoom) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnitBindRoom) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnitBindRoom) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCourseUnitBindRoom) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnitBindRoom) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnitBindRoom) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCourseUnitBindRoom) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnitBindRoom) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnitBindRoom) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmCourseUnitBindRoom) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "courseUnitId"
	eq1.Vy = bd.CourseUnitId
	req := request.Request{}
	req.Res = "BmCourseUnitBindRoom"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmCourseUnitBindRoom
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
