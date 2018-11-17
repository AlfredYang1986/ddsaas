package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionables struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Sessionables []BmSessionable `json:"Sessionables" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionables) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionables) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionables) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionables) QueryId() string {
	return bd.Id
}

func (bd *BmSessionables) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionables) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionables) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionables) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionables) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionables) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionables) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionables) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Sessionables)
	for i, r := range bd.Sessionables {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.Sessionables[i] = r
	}
	return err
}
