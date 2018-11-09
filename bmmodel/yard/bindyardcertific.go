package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindYardCertific struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	YardId string `json:"yardId" bson:"yardId"`
	CertificationId string `json:"certificationId" bson:"certificationId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindYardCertific) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindYardCertific) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindYardCertific) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindYardCertific) QueryId() string {
	return bd.Id
}

func (bd *BmBindYardCertific) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindYardCertific) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindYardCertific) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindYardCertific) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindYardCertific) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindYardCertific) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindYardCertific) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}