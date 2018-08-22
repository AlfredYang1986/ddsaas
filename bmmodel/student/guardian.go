package student

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMGuardian struct {
	Id        			string            	`json:"id"`
	Id_       			bson.ObjectId     	`bson:"_id"`

	Name      			string            	`json:"name" bson:"name"`
	NickName  			string           	`json:"nickname" bson:"nickname"`
	Relationship  		string           	`json:"relationship" bson:"relationship"`
	Mobile  			string           	`json:"mobile" bson:"mobile"`
	WeChatId  			string           	`json:"wechatid" bson:"wechatid"`
	Address  			string           	`json:"address" bson:"address"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMGuardian) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMGuardian) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMGuardian) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMGuardian) QueryId() string {
	return bd.Id
}

func (bd *BMGuardian) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMGuardian) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMGuardian) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMGuardian) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMGuardian) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMGuardian) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMGuardian) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}