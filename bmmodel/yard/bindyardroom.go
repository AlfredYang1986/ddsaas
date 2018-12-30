package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindYardRoom struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	YardId string `json:"yardId" bson:"yardId"`
	RoomId string `json:"roomId" bson:"roomId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindYardRoom) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindYardRoom) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindYardRoom) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindYardRoom) QueryId() string {
	return bd.Id
}

func (bd *BmBindYardRoom) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindYardRoom) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindYardRoom) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindYardRoom) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindYardRoom) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindYardRoom) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindYardRoom) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindYardRoom) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmBindYardRoom) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "yardId"
	eq1.Vy = bd.YardId
	eq2 := request.Eqcond{}
	eq2.Ky = "roomId"
	eq2.Vy = bd.RoomId
	req := request.Request{}
	req.Res = "BmBindYardRoom"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindYardRoom
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}

func (bd *BmBindYardRoom) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "yardId"
	eq1.Vy = bd.YardId
	req := request.Request{}
	req.Res = "BmBindYardRoom"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmBindYardRoom
	err := bind.DeleteAll(c.(request.Request))
	return err
}
