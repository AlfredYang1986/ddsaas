package honor

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmHonor struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Img string `json:"img" bson:"img"`
	Tag string `json:"tag" bson:"tag"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmHonor) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmHonor) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmHonor) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmHonor) QueryId() string {
	return bd.Id
}

func (bd *BmHonor) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmHonor) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmHonor) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmHonor) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmHonor) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmHonor) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmHonor) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
