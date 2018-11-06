package reward

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMReward struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	ImgSrc	  string 			`json:"img_src" bson:"img_src"`
	RewardDes	  string 			`json:"reward_des" bson:"reward_des"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMReward) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMReward) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMReward) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMReward) QueryId() string {
	return bd.Id
}

func (bd *BMReward) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMReward) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMReward) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMReward) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMReward) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMReward) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMReward) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
