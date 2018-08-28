package activity

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMActivityBrandRS struct {
	Id        		string            `json:"id"`
	Id_       		bson.ObjectId     `bson:"_id"`

	ActivityId	  	string 			`json:"activity_id" bson:"activity_id"`
	BrandId	  		string 			`json:"brand_id" bson:"brand_id"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMActivityBrandRS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMActivityBrandRS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMActivityBrandRS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMActivityBrandRS) QueryId() string {
	return bd.Id
}

func (bd *BMActivityBrandRS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMActivityBrandRS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMActivityBrandRS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMActivityBrandRS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMActivityBrandRS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMActivityBrandRS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMActivityBrandRS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}