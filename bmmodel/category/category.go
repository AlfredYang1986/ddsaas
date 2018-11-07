package category

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmCategory struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	//State  string `json:"state" bson:"state"`
	Title string `json:"title" bson:"title"`
	SubTitle string `json:"subtitle" bson:"subtitle"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCategory) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCategory) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCategory) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCategory) QueryId() string {
	return bd.Id
}

func (bd *BmCategory) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCategory) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCategory) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCategory) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCategory) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCategory) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCategory) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}