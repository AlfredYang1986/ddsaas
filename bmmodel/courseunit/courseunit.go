package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnit struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Status        float64                   `json:"status" bson:"status"`
	StartDate     float64                   `json:"start_date" bson:"start_date"`
	EndDate       float64                   `json:"end_date" bson:"end_date"`
	CourseTime    float64                   `json:"courseTime" bson:"courseTime"`		//课时
	SessionableId string                    `json:"sessionableId" bson:"sessionableId"`
	Sessionable   sessionable.BmSessionable `json:"Sessionable" jsonapi:"relationships"`
	Teacher       teacher.BmTeacher         `json:"Teacher" jsonapi:"relationships"`
	Room          room.BmRoom               `json:"Room" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnit) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnit) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnit) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnit) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnit) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Sessionable":
		bd.Sessionable = v.(sessionable.BmSessionable)
	case "Teacher":
		bd.Teacher = v.(teacher.BmTeacher)
	case "Room":
		bd.Room = v.(room.BmRoom)
	}
	return bd
}

func (bd BmCourseUnit) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnit) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCourseUnit) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnit) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnit) ReSetProp() error {
	bd.reSetSessionabel()
	bd.reSetTeacher()
	bd.reSetRoom()
	return nil
}

func (bd *BmCourseUnit) reSetSessionabel() error {

	eq := request.Eqcond{}
	eq.Ky = "courseUnitId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmCourseUnitBindSessionable"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmCourseUnitBindSessionable{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.SessionableId
	req0 := request.Request{}
	req0.Res = "BmSessionable"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := sessionable.BmSessionable{}
	err = result.FindOne(c0.(request.Request))
	result.ReSetProp()
	bd.Sessionable = result

	return err
}

func (bd *BmCourseUnit) reSetTeacher() error {

	eq := request.Eqcond{}
	eq.Ky = "courseUnitId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmCourseUnitBindTeacher"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmCourseUnitBindTeacher{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.TeacherId
	req0 := request.Request{}
	req0.Res = "BmTeacher"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := teacher.BmTeacher{}
	err = result.FindOne(c0.(request.Request))
	bd.Teacher = result

	return err
}

func (bd *BmCourseUnit) reSetRoom() error {

	eq := request.Eqcond{}
	eq.Ky = "courseUnitId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmCourseUnitBindRoom"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmCourseUnitBindRoom{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.RoomId
	req0 := request.Request{}
	req0.Res = "BmRoom"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := room.BmRoom{}
	err = result.FindOne(c0.(request.Request))
	bd.Room = result

	return err
}
