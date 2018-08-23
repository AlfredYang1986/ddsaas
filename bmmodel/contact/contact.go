package contact

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/order"
	"gopkg.in/mgo.v2/bson"
)

type Contact struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`

	Orders []order.Order `json:"orders" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *Contact) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *Contact) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *Contact) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *Contact) QueryId() string {
	return bd.Id
}

func (bd *Contact) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *Contact) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd Contact) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "orders":
		var rst []order.Order
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(order.Order))
		}
		bd.Orders = rst
	}
	return bd
}

func (bd Contact) QueryConnect(tag string) interface{} {
	switch tag {
	case "orders":
		return bd.Orders
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *Contact) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *Contact) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *Contact) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
