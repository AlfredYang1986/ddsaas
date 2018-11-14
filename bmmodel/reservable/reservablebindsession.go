package reservable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmReservableBindSession struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ReservableId string `json:"reservableId" bson:"reservableId"`
	SessionId    string `json:"sessionId" bson:"sessionId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmReservableBindSession) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmReservableBindSession) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmReservableBindSession) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmReservableBindSession) QueryId() string {
	return bd.Id
}

func (bd *BmReservableBindSession) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmReservableBindSession) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmReservableBindSession) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmReservableBindSession) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmReservableBindSession) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmReservableBindSession) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmReservableBindSession) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmReservableBindSession) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmReservableBindSession) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "reservableId"
	eq1.Vy = bd.ReservableId
	req := request.Request{}
	req.Res = "BmReservableBindSession"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmReservableBindSession
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}