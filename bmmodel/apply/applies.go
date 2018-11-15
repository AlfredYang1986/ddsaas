package apply

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmApplies struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Applies []BmApply `json:"Applies" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApplies) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApplies) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApplies) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApplies) QueryId() string {
	return bd.Id
}

func (bd *BmApplies) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApplies) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApplies) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmApplies) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApplies) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApplies) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApplies) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmApplies) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Applies)
	for i, r := range bd.Applies {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.Applies[i] = r
	}
	return err
}
