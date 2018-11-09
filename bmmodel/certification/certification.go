package certification

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCertification struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Img string `json:"img" bson:"img"`
	Tag string `json:"tag" bson:"tag"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCertification) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCertification) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCertification) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCertification) QueryId() string {
	return bd.Id
}

func (bd *BmCertification) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCertification) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCertification) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCertification) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCertification) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCertification) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCertification) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
