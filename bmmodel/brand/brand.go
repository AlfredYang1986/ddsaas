package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/location"
	"gopkg.in/mgo.v2/bson"
)

type BMBrand struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`
	Name      string            `json:"name" bson:"name"`
	Slogan    string            `json:"slogan" bson:"slogan"`
	Highlight []string          `json:"highlights" bson:"heighlights"`
	About     string            `json:"about" bson:"about"`
	Awards    map[string]string `json:"awards"`
	Attends   map[string]string `json:"attends"`
	Qualifier map[string]string `json:"qualifier"`
	Found 	  int64 			`json:"found"`

	Locations []location.BMLocation `json:"locations" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMBrand) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMBrand) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMBrand) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMBrand) QueryId() string {
	return bd.Id
}

func (bd *BMBrand) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMBrand) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMBrand) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "locations":
		var rst []location.BMLocation
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(location.BMLocation))
		}
		bd.Locations = rst
	}
	return bd
}

func (bd BMBrand) QueryConnect(tag string) interface{} {
	switch tag {
	case "locations":
		return bd.Locations
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMBrand) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMBrand) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMBrand) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
