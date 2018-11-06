package room

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMRoom struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title    string `json:"title" bson:"title"`
	Capacity string `json:"capacity" bson:"capacity"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMRoom) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMRoom) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMRoom) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMRoom) QueryId() string {
	return bd.Id
}

func (bd *BMRoom) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMRoom) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMRoom) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMRoom) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMRoom) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMRoom) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMRoom) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
