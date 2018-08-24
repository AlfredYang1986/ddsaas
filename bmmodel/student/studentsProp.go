package student

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMStudentsProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	StudentsProp []BMStudentProp `json:"students" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMStudentsProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMStudentsProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMStudentsProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMStudentsProp) QueryId() string {
	return bd.Id
}

func (bd *BMStudentsProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMStudentsProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMStudentsProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMStudentsProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMStudentsProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMStudentsProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMStudentsProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
