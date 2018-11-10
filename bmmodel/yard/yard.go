package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"gopkg.in/mgo.v2/bson"
)

type BmYard struct {
	Id          string        `json:"id"`
	Id_         bson.ObjectId `bson:"_id"`
	BrandId     string        `json:"brandId" bson:"brandId"`
	Title       string        `json:"title" bson:"title"`
	Cover       string        `json:"cover" bson:"cover"`
	Description string        `json:"description" bson:"description"`
	Around      string        `json:"around" bson:"around"`

	//Address address.BmAddress `json:"address" bson:"relationships"`
	/**
	 * 在构建过程中，yard可能成为地址搜索的条件
	 */
	Province string `json:"province" bson:"province"`
	City     string `json:"city" bson:"city"`
	District string `json:"district" bson:"district"`
	Address   string `json:"address" bson:"address"`
	TrafficInfo   string `json:"traffic_info" bson:"traffic_info"`

	//RoomCount float64 `json:"room_count"`
	/**
	 * 在构建过程中，除了排课逻辑，不会通过query到Room
	 */
	Rooms   []room.BmRoom     `json:"Rooms" jsonapi:"relationships"`
	TagImgs []tagimg.BmTagImg `json:"Tagimgs" jsonapi:"relationships"`

	//TODO:20181109新增的
	Attribute string `json:"attribute" bson:"attribute"`
	Scenario  string `json:"scenario" bson:"scenario"`
	//TODO:Certifications合并成TagImgs,添加category做区分.
	//Certifications []certification.BmCertification `json:"Certifications" jsonapi:"relationships"`
	Facilities     []interface{}                   `json:"facilities" bson:"facilities"`
	Friendly       []interface{}                   `json:"friendly" bson:"friendly"`
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
	case "Rooms":
		var rst []room.BmRoom
		for _, item := range v.([]interface{}) {
			tmp := item.(room.BmRoom)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Rooms = rst
	case "Tagimgs":
		var rst []tagimg.BmTagImg
		for _, item := range v.([]interface{}) {
			tmp := item.(tagimg.BmTagImg)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.TagImgs = rst
	//case "Certifications":
	//	var rst []certification.BmCertification
	//	for _, item := range v.([]interface{}) {
	//		tmp := item.(certification.BmCertification)
	//		if len(tmp.Id) > 0 {
	//			rst = append(rst, tmp)
	//		}
	//	}
	//	bd.Certifications = rst
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
