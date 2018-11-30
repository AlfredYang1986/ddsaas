package sessioninfo

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindSessionImg struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	SessionId string `json:"sessionId" bson:"sessionId"`
	TagImgId string `json:"tagImgId" bson:"tagImgId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindSessionImg) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindSessionImg) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindSessionImg) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindSessionImg) QueryId() string {
	return bd.Id
}

func (bd *BmBindSessionImg) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindSessionImg) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindSessionImg) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindSessionImg) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindSessionImg) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindSessionImg) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindSessionImg) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindSessionImg) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmBindSessionImg) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionId"
	eq1.Vy = bd.SessionId
	eq2 := request.Eqcond{}
	eq2.Ky = "tagImgId"
	eq2.Vy = bd.TagImgId
	req := request.Request{}
	req.Res = "BmBindSessionImg"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindSessionImg
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}