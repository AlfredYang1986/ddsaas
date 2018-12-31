package room

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmRooms struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Rooms []BmRoom `json:"Rooms" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmRooms) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmRooms) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmRooms) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmRooms) QueryId() string {
	return bd.Id
}

func (bd *BmRooms) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmRooms) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmRooms) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmRooms) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmRooms) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmRooms) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmRooms) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmRooms) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Rooms)
	for i, r := range bd.Rooms {
		r.ResetIdWithId_()
		bd.Rooms[i] = r
	}
	return err
}
