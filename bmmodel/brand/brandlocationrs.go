package brand

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMBrandLocationRS struct {
	Id        		string            	`json:"id"`
	Id_       		bson.ObjectId     	`bson:"_id"`

	BrandId 		string        		`json:"brand_id" bson:"brand_id"`
	LocationId 		string        		`json:"location_id" bson:"location_id"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMBrandLocationRS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMBrandLocationRS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMBrandLocationRS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMBrandLocationRS) QueryId() string {
	return bd.Id
}

func (bd *BMBrandLocationRS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMBrandLocationRS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMBrandLocationRS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMBrandLocationRS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMBrandLocationRS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMBrandLocationRS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMBrandLocationRS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}