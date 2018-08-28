package brand

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMBrandCompanyRS struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	BrandId 		string        		`json:"brand_id" bson:"brand_id"`
	CompanyId 		string        		`json:"company_id" bson:"company_id"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMBrandCompanyRS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMBrandCompanyRS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMBrandCompanyRS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMBrandCompanyRS) QueryId() string {
	return bd.Id
}

func (bd *BMBrandCompanyRS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMBrandCompanyRS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMBrandCompanyRS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMBrandCompanyRS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMBrandCompanyRS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMBrandCompanyRS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMBrandCompanyRS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}