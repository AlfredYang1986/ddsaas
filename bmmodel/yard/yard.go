package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
)

type BmYard struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title       string `json:"title" bson:"title"`
	Cover       string `json:"cover" bson:"cover"`
	Description string `json:"description" bson:"description"`
	Around      string `json:"around" bson:"around"`
	Facilities  string `json:"facilities" bson:"facilities"`

	//Address address.BmAddress `json:"address" bson:"relationships"`
	/**
	 * 在构建过程中，yard可能成为地址搜索的条件
	 */
	Province string `json:"province" bson:"province"`
	City string `json:"city" bson:"city"`
	District string `json:"district" bson:"district"`
	Detail string `json:"detail" bson:"detail"`

	//RoomCount float64 `json:"room_count"`
	/**
	 * 在构建过程中，除了排课逻辑，不会通过query到Room
	 */
	Rooms []room.BmRoom       `json:"rooms" jsonapi:"relationships"`
	TagImgs []tagimg.BmTagImg `json:"tagimgs" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmYard) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmYard) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmYard) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmYard) QueryId() string {
	return bd.Id
}

func (bd *BmYard) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmYard) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmYard) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "rooms":
		var rst []room.BmRoom
		for _, item := range v.([]interface{}) {
			tmp := item.(room.BmRoom)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Rooms = rst
	case "tagimgs":
		var rst []tagimg.BmTagImg
		for _, item := range v.([]interface{}) {
			tmp := item.(tagimg.BmTagImg)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.TagImgs = rst
	}
	return bd
}

func (bd BmYard) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmYard) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmYard) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmYard) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
