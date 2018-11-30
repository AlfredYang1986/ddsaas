package sessioninfo

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmSessionBindCat struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	SessionId string `json:"sessionId" bson:"sessionId"`
	CategoryId string `json:"categoryId" bson:"categoryId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionBindCat) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionBindCat) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionBindCat) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionBindCat) QueryId() string {
	return bd.Id
}

func (bd *BmSessionBindCat) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionBindCat) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionBindCat) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionBindCat) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionBindCat) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionBindCat) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionBindCat) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionBindCat) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmSessionBindCat) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionId"
	eq1.Vy = bd.SessionId
	eq2 := request.Eqcond{}
	eq2.Ky = "categoryId"
	eq2.Vy = bd.CategoryId
	req := request.Request{}
	req.Res = "BmSessionBindCat"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmSessionBindCat
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}