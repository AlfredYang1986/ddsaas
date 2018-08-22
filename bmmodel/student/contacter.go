package student

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMContacter struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Name      			string            	`json:"name" bson:"name"`
	NickName  			string           	`json:"nickname" bson:"nickname"`
	Relationship  		string           	`json:"relationship" bson:"relationship"`
	Mobile  			string           	`json:"mobile" bson:"mobile"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMContacter) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMContacter) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMContacter) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMContacter) QueryId() string {
	return bd.Id
}

func (bd *BMContacter) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMContacter) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMContacter) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMContacter) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMContacter) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMContacter) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMContacter) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}