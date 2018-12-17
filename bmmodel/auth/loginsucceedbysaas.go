package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"gopkg.in/mgo.v2/bson"
)

type BmLoginSucceedBySaaS struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Account  account.BmAccount  `json:"Account" jsonapi:"relationships"`
	Token string
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmLoginSucceedBySaaS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmLoginSucceedBySaaS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmLoginSucceedBySaaS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmLoginSucceedBySaaS) QueryId() string {
	return bd.Id
}

func (bd *BmLoginSucceedBySaaS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmLoginSucceedBySaaS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmLoginSucceedBySaaS) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Account":
		bd.Account = v.(account.BmAccount)
	}
	return bd
}

func (bd BmLoginSucceedBySaaS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmLoginSucceedBySaaS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmLoginSucceedBySaaS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmLoginSucceedBySaaS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
