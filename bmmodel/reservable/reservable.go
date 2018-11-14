package reservable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"gopkg.in/mgo.v2/bson"
)

type BmReservable struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Status    float64 `json:"status" bson:"status"` //0普通课程 1体验课 2活动
	StartDate float64 `json:"start_date" bson:"start_date"`
	EndDate   float64 `json:"end_date" bson:"end_date"`

	Yards       []yard.BmYard             `json:"Yards" jsonapi:"relationships"`
	SessionInfo sessioninfo.BmSessionInfo `json:"SessionInfo" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmReservable) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmReservable) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmReservable) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmReservable) QueryId() string {
	return bd.Id
}

func (bd *BmReservable) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmReservable) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmReservable) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Yards":
		var rst []yard.BmYard
		for _, item := range v.([]interface{}) {
			tmp := item.(yard.BmYard)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Yards = rst
	case "SessionInfo":
		bd.SessionInfo = v.(sessioninfo.BmSessionInfo)
	}
	return bd
}

func (bd BmReservable) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmReservable) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmReservable) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmReservable) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmReservable) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmReservable) ReSetProp() error {

	bd.reSetSessionInfo()
	bd.reSetYard()

	return nil
}

func (bd *BmReservable) reSetSessionInfo() error {

	eq := request.Eqcond{}
	eq.Ky = "reservableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmReservableBindSession"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmReservableBindSession{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.SessionId
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

func (bd *BmReservable) reSetYard() error {

	req := request.Request{}
	req.Res = "BmReservableBindYard"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "reservableId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmReservableBindYard
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.YardId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []yard.BmYard
	err = bmmodel.FindMutilWithBson("BmYard", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		ir.ReSetProp()
		resultArr[i] = ir
	}
	bd.Yards = resultArr

	return nil
}
