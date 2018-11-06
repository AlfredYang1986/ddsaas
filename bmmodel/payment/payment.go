package payment

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/transation"
	"gopkg.in/mgo.v2/bson"
)

type BMPayment struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Amount	  string 			`json:"amount" bson:"amount"`
	Description	  string 			`json:"description" bson:"description"`
	Date	  string 			`json:"date" bson:"date"`
	Status	  string 			`json:"status" bson:"status"`

	Transation []transation.BMTransation	`json:"transation" bson:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMPayment) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMPayment) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMPayment) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMPayment) QueryId() string {
	return bd.Id
}

func (bd *BMPayment) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMPayment) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMPayment) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMPayment) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMPayment) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMPayment) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMPayment) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
