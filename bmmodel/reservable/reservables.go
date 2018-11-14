package reservable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmReservables struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Reservables []BmReservable `json:"Reservables" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmReservables) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmReservables) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmReservables) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmReservables) QueryId() string {
	return bd.Id
}

func (bd *BmReservables) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmReservables) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmReservables) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmReservables) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmReservables) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmReservables) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmReservables) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmReservables) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Reservables)
	for i, r := range bd.Reservables {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.Reservables[i] = r
	}
	return err
}
