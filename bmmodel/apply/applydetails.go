package apply

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmApplyDetails struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Applies []BmApplyDetail `json:"Applies" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApplyDetails) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApplyDetails) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApplyDetails) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApplyDetails) QueryId() string {
	return bd.Id
}

func (bd *BmApplyDetails) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApplyDetails) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApplyDetails) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmApplyDetails) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApplyDetails) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApplyDetails) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApplyDetails) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmApplyDetails) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Applies)
	for i, r := range bd.Applies {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.Applies[i] = r
	}
	return err
}
