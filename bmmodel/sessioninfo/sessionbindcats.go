package sessioninfo

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmSessionInfoBindCats struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Binds []BmSessionBindCat `json:"binds" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionInfoBindCats) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionInfoBindCats) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionInfoBindCats) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionInfoBindCats) QueryId() string {
	return bd.Id
}

func (bd *BmSessionInfoBindCats) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionInfoBindCats) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionInfoBindCats) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "binds":
		var rst []BmSessionBindCat
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BmSessionBindCat))
		}
		bd.Binds = rst
	}
	return bd
}

func (bd BmSessionInfoBindCats) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionInfoBindCats) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionInfoBindCats) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionInfoBindCats) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionInfoBindCats) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Binds)
	for i, r := range bd.Binds {
		r.ResetIdWithId_()
		bd.Binds[i] = r
	}
	return err
}