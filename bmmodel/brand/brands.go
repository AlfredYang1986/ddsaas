package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBrands struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Brands []BmBrand `json:"Brands" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBrands) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBrands) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBrands) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBrands) QueryId() string {
	return bd.Id
}

func (bd *BmBrands) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBrands) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBrands) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBrands) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBrands) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBrands) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBrands) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBrands) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Brands)
	for i, r := range bd.Brands {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.Brands[i] = r
	}
	return err
}
