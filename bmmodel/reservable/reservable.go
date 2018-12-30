package reservable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"gopkg.in/mgo.v2/bson"
)

type BmReservable struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	BrandId    string  `json:"brandId" bson:"brandId"`
	Status     float64 `json:"status" bson:"status"` //0活动 1体验课 2普通课程
	StartDate  float64 `json:"start_date" bson:"start_date"`
	EndDate    float64 `json:"end_date" bson:"end_date"`
	CreateTime float64   `json:"create_time" bson:"create_time"`

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

func (bd *BmReservable) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmReservable) ReSetProp() error {
	bd.reSetSessionInfo()
	return nil
}

func (bd *BmReservable) DeleteProp() error {
	bd.deleteBmReservableBindSession()
	bd.deleteBmSessionable()
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

func (bd *BmReservable) deleteBmReservableBindSession() error {

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
	err = reval.DeleteAll(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.SessionId

	if eq0.Vy == "" {
		return err
	}

	req0 := request.Request{}
	req0.Res = "BmSessionInfo"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := sessioninfo.BmSessionInfo{}
	err = result.FindOne(c0.(request.Request))
	err = result.DeleteAll(c0.(request.Request))
	result.DeleteProp()
	bd.SessionInfo = result

	return err
}

func (bd *BmReservable) deleteBmSessionable() error {

	eq := request.Eqcond{}
	eq.Ky = "reservableId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmSessionable"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := sessionable.BmSessionable{}
	err := reval.DeleteAll(c.(request.Request))
	return err
}
