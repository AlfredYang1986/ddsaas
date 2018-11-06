package transation

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMTransation struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Date	  string 			`json:"date" bson:"date"`
	Method	  string 			`json:"method" bson:"method"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMTransation) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMTransation) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMTransation) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMTransation) QueryId() string {
	return bd.Id
}

func (bd *BMTransation) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMTransation) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMTransation) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMTransation) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMTransation) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMTransation) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMTransation) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
