package region

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMRegion struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Provice  string `json:"provice" bson:"provice"`
	City     string `json:"city" bson:"city"`
	District string `json:"district" bson:"district"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMRegion) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMRegion) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMRegion) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMRegion) QueryId() string {
	return bd.Id
}

func (bd *BMRegion) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMRegion) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMRegion) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMRegion) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMRegion) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMRegion) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMRegion) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
