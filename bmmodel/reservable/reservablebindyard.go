package reservable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmReservableBindYard struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ReservableId string `json:"reservableId" bson:"reservableId"`
	YardId       string `json:"yardId" bson:"yardId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmReservableBindYard) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmReservableBindYard) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmReservableBindYard) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmReservableBindYard) QueryId() string {
	return bd.Id
}

func (bd *BmReservableBindYard) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmReservableBindYard) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmReservableBindYard) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmReservableBindYard) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmReservableBindYard) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmReservableBindYard) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmReservableBindYard) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmReservableBindYard) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmReservableBindYard) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req, bd)
}

func (bd *BmReservableBindYard) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "reservableId"
	eq1.Vy = bd.ReservableId
	req := request.Request{}
	req.Res = "BmReservableBindYard"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmReservableBindYard
	err := bind.DeleteAll(c.(request.Request))
	return err
}