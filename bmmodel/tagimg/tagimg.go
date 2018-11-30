package tagimg

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmTagImg struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Img	  string 			`json:"img" bson:"img"`
	Tag	  string 			`json:"tag" bson:"tag"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmTagImg) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmTagImg) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmTagImg) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmTagImg) QueryId() string {
	return bd.Id
}

func (bd *BmTagImg) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmTagImg) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmTagImg) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmTagImg) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmTagImg) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmTagImg) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmTagImg) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmTagImg) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}
