package address

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/region"
	"gopkg.in/mgo.v2/bson"
)

type BmAddress struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Region region.BmRegion `json:"Region" jsonapi:"relationships"`
	Detail string          `json:"detail" bson:"detail"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmAddress) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmAddress) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmAddress) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmAddress) QueryId() string {
	return bd.Id
}

func (bd *BmAddress) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmAddress) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmAddress) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "region":
		bd.Region = v.(region.BmRegion)
	}
	return bd
}

func (bd BmAddress) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmAddress) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmAddress) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmAddress) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
