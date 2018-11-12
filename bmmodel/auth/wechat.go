package auth

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type BmWeChat struct {
	Id      string        `json:"id"`
	Id_     bson.ObjectId `bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Photo   string        `json:"photo" bson:"photo"`
	Open_id string        `json:"open_id" bson:"open_id"`
	Token string          `json:"token" bson:"token"`

	//TODO: 其它微信信息
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmWeChat) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmWeChat) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmWeChat) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmWeChat) QueryId() string {
	return bd.Id
}

func (bd *BmWeChat) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmWeChat) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmWeChat) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmWeChat) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmWeChat) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmWeChat) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmWeChat) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * wechat interface
 *------------------------------------------------*/

func (bd BmWeChat) IsWechatRegisted() bool {

	var once sync.Once
	var bmMongo bmconfig.BMMongoConfig
	once.Do(bmMongo.GenerateConfig)
	host := bmMongo.Host
	port := bmMongo.Port
	dbName := bmMongo.Database

	colName := "BmWeChat"

	session, err := mgo.Dial(host + ":" + port)
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB(dbName).C(colName)
	n, err := c.Find(bson.M{"open_id": bd.Open_id}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BmWeChat) Valid() bool {
	return bd.Open_id != ""
}
