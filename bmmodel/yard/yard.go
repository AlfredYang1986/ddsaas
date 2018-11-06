package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/address"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"gopkg.in/mgo.v2/bson"
)

type BMYard struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title       string `json:"title" bson:"title"`
	Cover       string `json:"cover" bson:"cover"`
	Description string `json:"description" bson:"description"`
	Around      string `json:"around" bson:"around"`
	Facilities  string `json:"facilities" bson:"facilities"`

	Address address.BMAddress `json:"address" bson:"relationships"`
	Rooms []room.BMRoom `json:"rooms" bson:"relationships"`
	TagImgs []tagimg.BMTagImg `json:"tagimgs" bson:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMYard) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMYard) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMYard) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMYard) QueryId() string {
	return bd.Id
}

func (bd *BMYard) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMYard) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMYard) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMYard) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMYard) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMYard) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMYard) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
