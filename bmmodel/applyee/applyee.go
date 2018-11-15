package applyee

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type BmApplyee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name          string              `json:"name" bson:"name"`
	Pic           string              `json:"pic" bson:"pic"`
	RegisterPhone string              `json:"regi_phone" bson:"regi_phone"`
	WeChatOpenid  string              `json:"wechat_openid" bson:"wechat_openid"`
	WeChatToken   string              `json:"wechat_token" bson:"wechat_token"`
}

/*-----------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmApplyee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmApplyee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmApplyee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmApplyee) QueryId() string {
	return bd.Id
}

func (bd *BmApplyee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmApplyee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmApplyee) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmApplyee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmApplyee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmApplyee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmApplyee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

var once sync.Once
var bmMongoConfig bmconfig.BMMongoConfig

func (bd BmApplyee) IsRegisted() bool {
	once.Do(bmMongoConfig.GenerateConfig)
	session, err := mgo.Dial(bmMongoConfig.Host + ":" + bmMongoConfig.Port)
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB(bmMongoConfig.Database).C("BmApplyee")
	n, err := c.Find(bson.M{"wechat_openid": bd.WeChatOpenid}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BmApplyee) Valid() bool {
	return bd.WeChatOpenid != ""
}
