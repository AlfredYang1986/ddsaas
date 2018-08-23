package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMTeacher struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name     string `json:"name" bson:"name"`
	NickName string `json:"nickname" bson:"nickname"`
	Birthday string `json:"birthday" bson:"birthday"`
	Age      int    `json:"age" bson:"age"`
	Sex      string `json:"sex" bson:"sex"`
	City     string `json:"city" bson:"city"`
	Mobile   string `json:"mobile" bson:"mobile"`
	WeChatId string `json:"wechatid" bson:"wechatid"`
	Address  string `json:"address" bson:"address"`
	Photo    string `json:"photo" bson:"photo"`

	WorkExperience      []interface{} `json:"workexperience" bson:"workexperience"`
	EducationExperience []interface{} `json:"educationexperience" bson:"educationexperience"`

	Found int64 `json:"found" bson:"found"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMTeacher) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMTeacher) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMTeacher) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMTeacher) QueryId() string {
	return bd.Id
}

func (bd *BMTeacher) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMTeacher) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMTeacher) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMTeacher) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMTeacher) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMTeacher) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMTeacher) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
