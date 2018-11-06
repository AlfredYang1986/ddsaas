package student

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"gopkg.in/mgo.v2/bson"
)

type BMStudent struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Person person.BmPerson          `json:"person" jsonapi:"relationships"`
	Guardians []guardian.BmGuardian `json:"guardians" jsonapi:"relationships"`
	//Contacts  []BMContacter `json:"contacts" jsonapi:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMStudent) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMStudent) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMStudent) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMStudent) QueryId() string {
	return bd.Id
}

func (bd *BMStudent) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMStudent) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMStudent) SetConnect(tag string, v interface{}) interface{} {
	//switch tag {
	//case "guardians":
	//	var guardians []BmGuardian
	//	for _, item := range v.([]interface{}) {
	//		guardians = append(guardians, item.(BmGuardian))
	//	}
	//	bd.Guardians = guardians
	//}
	return bd
}

func (bd BMStudent) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMStudent) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMStudent) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMStudent) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
