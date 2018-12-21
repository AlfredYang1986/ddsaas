package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/certification"
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
	Province       string        `json:"province" bson:"province"`
	City           string        `json:"city" bson:"city"`
	District       string        `json:"district" bson:"district"`
	Address        string        `json:"address" bson:"address"`
	TrafficInfo    string        `json:"traffic_info" bson:"traffic_info"`
	Attribute      string        `json:"attribute" bson:"attribute"`
	Scenario       string        `json:"scenario" bson:"scenario"`
	OpenTime       string        `json:"openTime" bson:"openTime"`
	ServiceContact string        `json:"serviceContact" bson:"serviceContact"`
	Facilities     []interface{} `json:"facilities" bson:"facilities"`
	//Friendly       []interface{}                   `json:"friendly" bson:"friendly"`

	//RoomCount float64 `json:"room_count"`
	/**
	 * 在构建过程中，除了排课逻辑，不会通过query到Room
	 */
	//TODO:Certifications合并成TagImgs,添加category做区分.
	Rooms          []room.BmRoom                   `json:"Rooms" jsonapi:"relationships"`
	TagImgs        []tagimg.BmTagImg               `json:"Tagimgs" jsonapi:"relationships"`
	Certifications []certification.BmCertification `json:"Certifications" jsonapi:"relationships"`
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
	case "Certifications":
		var rst []certification.BmCertification
		for _, item := range v.([]interface{}) {
			tmp := item.(certification.BmCertification)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.Certifications = rst
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

func (bd *BmYard) ReSetProp() error {

	bd.reSetRooms()
	bd.reSetTagImg()
	bd.reSetCertifications()

	return nil
}

func (bd *BmYard) reSetRooms() error {

	req := request.Request{}
	req.Res = "BmBindYardRoom"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "yardId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindYardRoom
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.RoomId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var rooms []room.BmRoom
	err = bmmodel.FindMutilWithBson("BmRoom", or_condi, &rooms)

	for i, ir := range rooms {
		ir.ResetIdWithId_()
		rooms[i] = ir
	}
	bd.Rooms = rooms

	return err
}

func (bd *BmYard) reSetTagImg() error {

	req := request.Request{}
	req.Res = "BmBindYardImg"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "yardId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindYardImg
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.TagImgId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var imgs []tagimg.BmTagImg
	err = bmmodel.FindMutilWithBson("BmTagImg", or_condi, &imgs)

	for i, ir := range imgs {
		ir.ResetIdWithId_()
		imgs[i] = ir
	}
	bd.TagImgs = imgs

	return err
}

func (bd *BmYard) reSetCertifications() error {

	req := request.Request{}
	req.Res = "BmBindYardCertific"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "yardId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindYardCertific
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.CertificationId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var cert []certification.BmCertification
	err = bmmodel.FindMutilWithBson("BmCertification", or_condi, &cert)

	for i, ir := range cert {
		ir.ResetIdWithId_()
		cert[i] = ir
	}
	bd.Certifications = cert

	return err
}
