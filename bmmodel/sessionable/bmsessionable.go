package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/count"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BmSessionable struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ClassTitle        string  `json:"classTitle" bson:"classTitle"`
	Status            float64 `json:"status" bson:"status"` //0活动 1体验课 2普通课程
	StartDate         float64 `json:"start_date" bson:"start_date"`
	EndDate           float64 `json:"end_date" bson:"end_date"`
	CreateTime        float64 `json:"create_time" bson:"create_time"`
	CourseTotalCount  float64 `json:"courseTotalCount"`
	CourseExpireCount float64 `json:"courseExpireCount"`
	BrandId           string  `json:"brandId" bson:"brandId"`
	ReservableId      string  `json:"reservableId" bson:"reservableId"`

	Yard        yard.BmYard               `json:"Yard" jsonapi:"relationships"`
	SessionInfo sessioninfo.BmSessionInfo `json:"SessionInfo" jsonapi:"relationships"`
	Teachers    []teacher.BmTeacher       `json:"Teachers" jsonapi:"relationships"`
	Attendees   []attendee.BmAttendee     `json:"Attendees" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionable) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionable) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionable) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionable) QueryId() string {
	return bd.Id
}

func (bd *BmSessionable) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionable) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionable) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Teachers":
		var rst []teacher.BmTeacher
		for _, item := range v.([]interface{}) {
			tmp := item.(teacher.BmTeacher)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Teachers = rst
	case "Attendees":
		var rst []attendee.BmAttendee
		for _, item := range v.([]interface{}) {
			tmp := item.(attendee.BmAttendee)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Attendees = rst
	case "Yard":
		bd.Yard = v.(yard.BmYard)
	case "SessionInfo":
		bd.SessionInfo = v.(sessioninfo.BmSessionInfo)
	}
	return bd
}

func (bd BmSessionable) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionable) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionable) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmSessionable) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionable) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionable) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmSessionable) ReSetProp() error {
	bd.reSetYard()
	bd.reSetSessionInfo()
	bd.reSetTeachers()
	bd.reSetAttendees()
	bd.reSetCourseCount()
	return nil
}

func (bd *BmSessionable) DeleteProp() error {
	bd.deleteBmSessionableBindYard()
	bd.deleteBmSessionableBindSessionInfo()
	bd.deleteBmSessionableBindAttendee()
	bd.deleteBmSessionableBindTeacher()
	return nil
}

func (bd *BmSessionable) reSetYard() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionableBindYard"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionableBindYard{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.YardId
	req0 := request.Request{}
	req0.Res = "BmYard"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := yard.BmYard{}
	err = result.FindOne(c0.(request.Request))
	result.ReSetProp()
	bd.Yard = result

	return err
}

func (bd *BmSessionable) reSetSessionInfo() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionableBindSessionInfo"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionableBindSessionInfo{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.SessionInfoId
	req0 := request.Request{}
	req0.Res = "BmSessionInfo"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := sessioninfo.BmSessionInfo{}
	err = result.FindOne(c0.(request.Request))
	result.ReSetProp()
	bd.SessionInfo = result

	return err
}

func (bd *BmSessionable) reSetTeachers() error {

	req := request.Request{}
	req.Res = "BmSessionableBindTeacher"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmSessionableBindTeacher
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.TeacherId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []teacher.BmTeacher
	err = bmmodel.FindMutilWithBson("BmTeacher", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		resultArr[i] = ir
	}
	bd.Teachers = resultArr

	return nil
}

func (bd *BmSessionable) reSetAttendees() error {

	req := request.Request{}
	req.Res = "BmSessionableBindAttendee"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmSessionableBindAttendee
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.AttendeeId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []attendee.BmAttendee
	err = bmmodel.FindMutilWithBson("BmAttendee", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		ir.ReSetProp()
		resultArr[i] = ir
	}
	bd.Attendees = resultArr

	return nil
}

func (bd *BmSessionable) reSetCourseCount() error {

	req := request.Request{}
	req.Res = "BmCourseUnit"
	var eqCondi []interface{}
	var ltCondi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	eqCondi = append(eqCondi, eq)
	c := req.SetConnect("conditions", eqCondi)
	req1 := c.(request.Request)
	var tmp count.BmCount
	totalCount, err := tmp.FindCount(req1)
	bd.CourseTotalCount = float64(totalCount)

	lt := request.Ltcond{}
	lt.Ky = "end_date"
	lt.Vy = float64(time.Now().UnixNano() / 1e6)
	ltCondi = append(ltCondi, lt)
	ltc := req1.SetConnect("Ltcond", ltCondi)
	expireCount, err := tmp.FindCount(ltc.(request.Request))
	bd.CourseExpireCount = float64(expireCount)

	return err
}

func (bd *BmSessionable) deleteBmSessionableBindYard() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionableBindYard"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionableBindYard{}
	err := reval.DeleteAll(c.(request.Request))

	return err
}

func (bd *BmSessionable) deleteBmSessionableBindSessionInfo() error {

	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionableBindSessionInfo"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmSessionableBindSessionInfo{}
	err := reval.DeleteAll(c.(request.Request))

	return err
}

func (bd *BmSessionable) deleteBmSessionableBindTeacher() error {

	req := request.Request{}
	req.Res = "BmSessionableBindTeacher"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var bind BmSessionableBindTeacher
	err := bind.DeleteAll(c.(request.Request))

	return err
}

func (bd *BmSessionable) deleteBmSessionableBindAttendee() error {

	req := request.Request{}
	req.Res = "BmSessionableBindAttendee"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var bind BmSessionableBindAttendee
	err := bind.DeleteAll(c.(request.Request))

	return err
}
