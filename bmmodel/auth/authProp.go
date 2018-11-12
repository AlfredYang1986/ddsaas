package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmAuthProp struct {
	Id         string        `json:"Id"`
	Id_        bson.ObjectId `bson:"_id"`
	Auth_id    string        `json:"auth_id" bson:"auth_id"`
	Phone_id   string        `json:"phone_id" bson:"phone_id"`
	Wechat_id  string        `json:"wechat_id" bson:"wechat_id"`
	Profile_id string        `json:"profile_id" bson:"profile_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmAuthProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmAuthProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmAuthProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmAuthProp) QueryId() string {
	return bd.Id
}

func (bd *BmAuthProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmAuthProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmAuthProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmAuthProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmAuthProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmAuthProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAuthProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
