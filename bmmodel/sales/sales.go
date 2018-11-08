package sales

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMSales struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	UniqueId	  string   `json:"uniqueId" bson:"uniqueId"`
	//Person person.BmPerson `json:"person" jsonapi:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMSales) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMSales) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMSales) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMSales) QueryId() string {
	return bd.Id
}

func (bd *BMSales) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMSales) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMSales) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMSales) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMSales) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMSales) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMSales) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
