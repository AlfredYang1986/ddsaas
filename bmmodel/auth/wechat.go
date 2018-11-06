package auth

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type BMWeChat struct {
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

func (bd *BMWeChat) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMWeChat) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMWeChat) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMWeChat) QueryId() string {
	return bd.Id
}

func (bd *BMWeChat) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMWeChat) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMWeChat) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMWeChat) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMWeChat) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMWeChat) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMWeChat) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * wechat interface
 *------------------------------------------------*/

func (bd BMWeChat) IsWechatRegisted() bool {

	var once sync.Once
	var bmMongo bmconfig.BMMongoConfig
	once.Do(bmMongo.GenerateConfig)
	host := bmMongo.Host
	port := bmMongo.Port
	dbName := bmMongo.Database

	colName := "BMWeChat"

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

func (bd BMWeChat) Valid() bool {
	return bd.Open_id != ""
}
