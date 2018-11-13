package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmAuth struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Phone  BmPhone  `json:"BmPhone" jsonapi:"relationships"`
	Wechat BmWeChat `json:"BmWeChat" jsonapi:"relationships"`

	Token string
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmAuth) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmAuth) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmAuth) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmAuth) QueryId() string {
	return bd.Id
}

func (bd *BmAuth) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmAuth) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmAuth) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "BmPhone":
		bd.Phone = v.(BmPhone)
	case "BmWeChat":
		bd.Wechat = v.(BmWeChat)
	//case "profile":
	//	bd.Profile = v.(profile.BMProfile)
	}
	return bd
}

func (bd BmAuth) QueryConnect(tag string) interface{} {
	switch tag {
	case "BmPhone":
		return bd.Phone
	case "BmWeChat":
		return bd.Wechat
	//case "profile":
	//	return bd.Profile
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmAuth) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmAuth) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAuth) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
