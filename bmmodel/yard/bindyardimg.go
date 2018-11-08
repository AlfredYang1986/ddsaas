package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindYardImg struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	YardId string `json:"yardId" bson:"yardId"`
	TagImgId string `json:"tagImgId" bson:"tagImgId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindYardImg) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindYardImg) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindYardImg) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindYardImg) QueryId() string {
	return bd.Id
}

func (bd *BmBindYardImg) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindYardImg) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindYardImg) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindYardImg) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindYardImg) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindYardImg) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindYardImg) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}