package region

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmRegion struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Provice  string `json:"provice" bson:"provice"`
	City     string `json:"city" bson:"city"`
	District string `json:"district" bson:"district"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmRegion) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmRegion) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmRegion) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmRegion) QueryId() string {
	return bd.Id
}

func (bd *BmRegion) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmRegion) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmRegion) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmRegion) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmRegion) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmRegion) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmRegion) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
