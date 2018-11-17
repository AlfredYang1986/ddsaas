package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionableBindSessionInfo struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	SessionableId string `json:"sessionableId" bson:"sessionableId"`
	SessionInfoId        string `json:"sessionInfoId" bson:"sessionInfoId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionableBindSessionInfo) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionableBindSessionInfo) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionableBindSessionInfo) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionableBindSessionInfo) QueryId() string {
	return bd.Id
}

func (bd *BmSessionableBindSessionInfo) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionableBindSessionInfo) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionableBindSessionInfo) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionableBindSessionInfo) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionableBindSessionInfo) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionableBindSessionInfo) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmSessionableBindSessionInfo) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionableBindSessionInfo) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionableBindSessionInfo) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionableId"
	eq1.Vy = bd.SessionableId
	req := request.Request{}
	req.Res = "BmSessionableBindSessionInfo"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmSessionableBindSessionInfo
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
