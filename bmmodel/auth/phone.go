package auth

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type BMPhone struct {
	Id      string        `json:"id"`
	Id_     bson.ObjectId `bson:"_id"`
	PhoneNo string        `json:"phone_no" bson:"phone_no"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMPhone) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMPhone) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMPhone) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMPhone) QueryId() string {
	return bd.Id
}

func (bd *BMPhone) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMPhone) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMPhone) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMPhone) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMPhone) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMPhone) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMPhone) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * phone interface
 *------------------------------------------------*/

func (bd BMPhone) IsPhoneRegisted() bool {

	var once sync.Once
	var bmMongo bmconfig.BMMongoConfig
	once.Do(bmMongo.GenerateConfig)
	host := bmMongo.Host
	port := bmMongo.Port
	dbName := bmMongo.Database

	colName := "BMPhone"

	session, err := mgo.Dial(host + ":" + port)
	if err != nil {
		panic("dial dbName error")
	}
	defer session.Close()

	c := session.DB(dbName).C(colName)
	n, err := c.Find(bson.M{"phone_no": bd.PhoneNo}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMPhone) Valid() bool {
	return bd.PhoneNo != ""
}
