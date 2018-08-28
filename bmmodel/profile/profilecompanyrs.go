package profile

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMProfileCompanyRS struct {
	Id        	string            	`json:"id"`
	Id_       	bson.ObjectId     	`bson:"_id"`

	ProfileId 	string 				`json:"profile_id" bson:"profile_id"`
	CompanyId 	string 				`json:"company_id" bson:"company_id"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMProfileCompanyRS) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMProfileCompanyRS) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMProfileCompanyRS) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMProfileCompanyRS) QueryId() string {
	return bd.Id
}

func (bd *BMProfileCompanyRS) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMProfileCompanyRS) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMProfileCompanyRS) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMProfileCompanyRS) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMProfileCompanyRS) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMProfileCompanyRS) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMProfileCompanyRS) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}