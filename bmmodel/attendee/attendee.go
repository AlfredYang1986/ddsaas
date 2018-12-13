package attendee

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/payment"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type BmAttendee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	BrandId string `json:"brandId" bson:"brandId"`

	KidId   string `json:"kidId" bson:"kidId"`
	ApplyId string `json:"applyId" bson:"applyId"`

	TeacherId   string `json:"teacherId" bson:"teacherId"`
	TeacherName string `json:"teacherName" bson:"teacherName"`
	SourceWay   string `json:"sourceWay" bson:"sourceWay"`

	Intro       string  `json:"intro" bson:"intro"`
	Status      string  `json:"status" bson:"status"` //未付款-candidate, 已付款-stud
	LessonCount float64 `json:"lesson_count" bson:"lesson_count"`

	Name     string  `json:"name" bson:"name"`
	Nickname string  `json:"nickname" bson:"nickname"`
	Icon     string  `json:"icon" bson:"icon"`
	Dob      float64 `json:"dob" bson:"dob"`
	Gender   float64 `json:"gender" bson:"gender"`
	RegDate  float64 `json:"reg_date" bson:"reg_date"`
	Contact  string  `json:"contact" bson:"contact"`
	WeChat   string  `json:"wechat" bson:"wechat"`

	//Person    person.BmPerson       `json:"Person" jsonapi:"relationships"`
	Guardians []guardian.BmGuardian `json:"Guardians" jsonapi:"relationships"`
	Applyees  []applyee.BmApplyee   `json:"Applyees" jsonapi:"relationships"`
	Payments  []payment.BMPayment   `json:"Payments" jsonapi:"relationships"`
	Address   string                `json:"address" bson:"address"`
	School    string                `json:"school" bson:"school"`
	IdCardNo  string                `json:"idCardNo" bson:"idCardNo"`
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
	case "Applyees":
		var rst []applyee.BmApplyee
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(applyee.BmApplyee))
		}
		bd.Applyees = rst
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

func (bd *BmAttendee) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmAttendee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAttendee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmAttendee) GetAttendeeGuardianRSes() (error, []BMAttendeeGuardianRS) {

	eq2 := request.Eqcond{}
	var eq2arr []request.Eqcond
	eq2.Ky = "attendeeId"
	eq2.Vy = bd.Id
	eq2.Ct = "BMAttendeeGuardianRS"
	req2 := request.Request{}
	req2.Res = "BMAttendeeGuardianRS"
	req2.Eqcond = append(eq2arr, eq2)
	var condi2 []interface{}
	condi2 = append(condi2, eq2)
	c2 := req2.SetConnect("Eqcond", condi2)
	agrsarr := BMAttendeeGuardianRSeS{}
	err := agrsarr.FindMulti(c2.(request.Request))
	return err, agrsarr.AgRsArr

}

var once sync.Once
var bmMongoConfig bmconfig.BMMongoConfig

func (bd BmAttendee) IsAttendeeExist() bool {

	once.Do(bmMongoConfig.GenerateConfig)
	session, err := mgo.Dial(bmMongoConfig.Host + ":" + bmMongoConfig.Port)

	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB(bmMongoConfig.Database).C("BmAttendee")
	n, err := c.Find(bson.M{"_id": bson.ObjectIdHex(bd.Id)}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd *BmAttendee) ReSetProp() error {
	bd.reSetGuardians()
	bd.reSetApplyees()
	return nil
}

func (bd *BmAttendee) reSetGuardians() error {

	req := request.Request{}
	req.Res = "BMAttendeeGuardianRS"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "attendeeId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BMAttendeeGuardianRS
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.GuardianId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []guardian.BmGuardian
	err = bmmodel.FindMutilWithBson("BmGuardian", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		resultArr[i] = ir
	}
	bd.Guardians = resultArr

	return nil
}

func (bd *BmAttendee) reSetApplyees() error {

	req := request.Request{}
	req.Res = "BMAttendeeBindApplyee"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "attendeeId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BMAttendeeBindApplyee
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.ApplyeeId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []applyee.BmApplyee
	err = bmmodel.FindMutilWithBson("BmApplyee", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		resultArr[i] = ir
	}
	bd.Applyees = resultArr

	return nil
}
