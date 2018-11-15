package apply

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmApplyBindKid struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ApplyId string `json:"applyId" bson:"applyId"`
	KidId   string `json:"kidId" bson:"kidId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApplyBindKid) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApplyBindKid) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApplyBindKid) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApplyBindKid) QueryId() string {
	return bd.Id
}

func (bd *BmApplyBindKid) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApplyBindKid) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApplyBindKid) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmApplyBindKid) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApplyBindKid) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApplyBindKid) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApplyBindKid) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmApplyBindKid) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "applyId"
	eq1.Vy = bd.ApplyId
	eq2 := request.Eqcond{}
	eq2.Ky = "kidId"
	eq2.Vy = bd.KidId
	req := request.Request{}
	req.Res = "BmApplyBindKid"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmApplyBindKid
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
