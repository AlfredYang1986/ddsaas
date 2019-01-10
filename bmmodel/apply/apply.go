package apply

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/ddsaas/bmmodel/kid"
	"gopkg.in/mgo.v2/bson"
)

type BmApply struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Status       float64           `json:"status" bson:"status"` //0=未处理，1=已处理
	ApplyTime    float64           `json:"apply_time" bson:"apply_time"`
	ExceptTime   float64           `json:"except_time" bson:"except_time"`
	CreateTime   float64           `json:"create_time" bson:"create_time"`
	ApplyeeId    string            `json:"applyeeId" bson:"applyeeId"`
	BrandId      string            `json:"brandId" bson:"brandId"`
	ApplyFrom    string            `json:"applyFrom" bson:"applyFrom"`
	CourseType   float64           `json:"courseType" bson:"courseType"` //0活动 1体验课 2普通课程 -1预注册
	CourseName   string            `json:"courseName" bson:"courseName"`
	Contact      string            `json:"contact" bson:"contact"`
	Kids         []kid.BmKid       `json:"Kids" jsonapi:"relationships"`
	Applyee      applyee.BmApplyee `json:"Applyee" jsonapi:"relationships"`
	ReservableId string            `json:"reservableId" bson:"reservableId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApply) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApply) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApply) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApply) QueryId() string {
	return bd.Id
}

func (bd *BmApply) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApply) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApply) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Applyee":
		bd.Applyee = v.(applyee.BmApplyee)
	case "Kids":
		var rst []kid.BmKid
		for _, item := range v.([]interface{}) {
			tmp := item.(kid.BmKid)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Kids = rst
	}
	return bd
}

func (bd BmApply) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApply) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApply) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmApply) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApply) FindCount(req request.Request) (int, error) {
	return bmmodel.FindCount(req)
}

func (bd *BmApply) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmApply) ReSetProp() error {

	bd.reSetApplyee()
	bd.reSetKids()

	return nil
}

func (bd *BmApply) reSetApplyee() error {

	eq := request.Eqcond{}
	eq.Ky = "applyId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmApplyBindApplyee"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmApplyBindApplyee{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.ApplyeeId
	req0 := request.Request{}
	req0.Res = "BmApplyee"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := applyee.BmApplyee{}
	err = result.FindOne(c0.(request.Request))
	bd.Applyee = result

	return err
}

func (bd *BmApply) reSetKids() error {

	req := request.Request{}
	req.Res = "BmApplyBindKid"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "applyId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmApplyBindKid
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.KidId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var results []kid.BmKid
	err = bmmodel.FindMutilWithBson("BmKid", or_condi, &results)

	for i, ir := range results {
		ir.ResetIdWithId_()
		results[i] = ir
	}
	bd.Kids = results

	return err
}
