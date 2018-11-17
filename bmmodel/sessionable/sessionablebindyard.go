package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionableBindYard struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	SessionableId string `json:"sessionableId" bson:"sessionableId"`
	YardId        string `json:"yardId" bson:"yardId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionableBindYard) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionableBindYard) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionableBindYard) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionableBindYard) QueryId() string {
	return bd.Id
}

func (bd *BmSessionableBindYard) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionableBindYard) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionableBindYard) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionableBindYard) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionableBindYard) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionableBindYard) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmSessionableBindYard) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionableBindYard) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionableBindYard) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionableId"
	eq1.Vy = bd.SessionableId
	req := request.Request{}
	req.Res = "BmSessionableBindYard"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmSessionableBindYard
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
