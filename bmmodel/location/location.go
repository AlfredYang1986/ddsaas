package location

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMLocation struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title         string        `json:"title" bson:"title"`
	Address       string        `json:"address" bson:"address"`
	Description   string        `json:"description" bson:"description"`
	Facilities    []interface{} `json:"facilities" bson:"facilities"`
	Environment   []interface{} `json:"environment" bson:"environment"`
	Certification []interface{} `json:"certification" bson:"certification"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (loc *BMLocation) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(loc)
}

func (loc *BMLocation) ResetId_WithID() {
	bmmodel.ResetId_WithID(loc)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMLocation) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMLocation) QueryId() string {
	return bd.Id
}

func (bd *BMLocation) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMLocation) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/

func (loc BMLocation) SetConnect(tag string, v interface{}) interface{} {
	return loc
}

func (loc BMLocation) QueryConnect(tag string) interface{} {
	return loc
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (loc *BMLocation) InsertBMObject() error {
	return bmmodel.InsertBMObject(loc)
}

func (loc *BMLocation) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, loc)
}

func (loc *BMLocation) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, loc)
}
