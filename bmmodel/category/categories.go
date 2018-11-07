package category

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmCategories struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Cats []BmCategory `json:"cats" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCategories) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCategories) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCategories) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCategories) QueryId() string {
	return bd.Id
}

func (bd *BmCategories) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCategories) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCategories) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "cats":
		var rst []BmCategory
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BmCategory))
		}
		bd.Cats = rst
	}
	return bd
}

func (bd BmCategories) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCategories) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCategories) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCategories) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCategories) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Cats)
	for i, r := range bd.Cats {
		r.ResetIdWithId_()
		bd.Cats[i] = r
	}
	return err
}