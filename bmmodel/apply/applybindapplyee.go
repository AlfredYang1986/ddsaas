package apply

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmApplyBindApplyee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ApplyId   string `json:"applyId" bson:"applyId"`
	ApplyeeId string `json:"applyeeId" bson:"applyeeId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApplyBindApplyee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApplyBindApplyee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApplyBindApplyee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApplyBindApplyee) QueryId() string {
	return bd.Id
}

func (bd *BmApplyBindApplyee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApplyBindApplyee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApplyBindApplyee) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmApplyBindApplyee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApplyBindApplyee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApplyBindApplyee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApplyBindApplyee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmApplyBindApplyee) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "applyId"
	eq1.Vy = bd.ApplyId
	eq2 := request.Eqcond{}
	eq2.Ky = "applyeeId"
	eq2.Vy = bd.ApplyeeId
	req := request.Request{}
	req.Res = "BmApplyBindApplyee"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmApplyBindApplyee
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
