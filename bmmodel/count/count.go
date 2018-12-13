package count

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCount struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Res   string  `json:"res" bson:"res"`
	Count int `json:"count" bson:"count"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCount) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCount) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCount) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCount) QueryId() string {
	return bd.Id
}

func (bd *BmCount) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCount) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCount) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCount) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCount) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCount) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCount) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCount) FindCount(req request.Request) (int, error) {
	return bmmodel.FindCount(req)
}

func (bd *BmCount) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
