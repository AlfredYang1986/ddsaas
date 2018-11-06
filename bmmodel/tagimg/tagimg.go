package tagimg

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMTagImg struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Img	  string 			`json:"img" bson:"img"`
	Tag	  string 			`json:"tag" bson:"tag"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMTagImg) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMTagImg) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMTagImg) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMTagImg) QueryId() string {
	return bd.Id
}

func (bd *BMTagImg) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMTagImg) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMTagImg) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMTagImg) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMTagImg) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMTagImg) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMTagImg) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
