package address

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/region"
	"gopkg.in/mgo.v2/bson"
)

type BMAddress struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Region region.BMRegion `json:"region" jsonapi:"relationships"`
	Detail string          `json:"detail" bson:"detail"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAddress) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAddress) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAddress) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAddress) QueryId() string {
	return bd.Id
}

func (bd *BMAddress) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAddress) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAddress) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAddress) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAddress) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAddress) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAddress) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
