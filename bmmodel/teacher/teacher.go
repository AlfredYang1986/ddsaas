package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMTeacher struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Person person.BmPerson `json:"person" jsonapi:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMTeacher) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMTeacher) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMTeacher) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMTeacher) QueryId() string {
	return bd.Id
}

func (bd *BMTeacher) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMTeacher) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMTeacher) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMTeacher) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMTeacher) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMTeacher) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMTeacher) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
