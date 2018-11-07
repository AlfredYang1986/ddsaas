package sessioninfo

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BmSessionInfos struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Sessions  []BmSessionInfo `json:"sessions" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionInfos) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionInfos) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionInfos) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionInfos) QueryId() string {
	return bd.Id
}

func (bd *BmSessionInfos) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionInfos) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionInfos) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "sessions":
		var rst []BmSessionInfo
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BmSessionInfo))
		}
		bd.Sessions = rst
	}
	return bd
}

func (bd BmSessionInfos) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionInfos) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionInfos) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionInfos) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionInfos) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Sessions)
	for i, r := range bd.Sessions {
		r.ResetIdWithId_()
		bd.Sessions[i] = r
	}
	return err
}