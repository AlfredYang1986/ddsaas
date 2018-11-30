package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionableBindAttendee struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	SessionableId string `json:"sessionableId" bson:"sessionableId"`
	AttendeeId    string `json:"attendeeId" bson:"attendeeId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionableBindAttendee) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionableBindAttendee) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionableBindAttendee) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionableBindAttendee) QueryId() string {
	return bd.Id
}

func (bd *BmSessionableBindAttendee) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionableBindAttendee) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionableBindAttendee) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionableBindAttendee) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionableBindAttendee) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionableBindAttendee) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmSessionableBindAttendee) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionableBindAttendee) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionableBindAttendee) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req)
}

func (bd *BmSessionableBindAttendee) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionableId"
	eq1.Vy = bd.SessionableId
	req := request.Request{}
	req.Res = "BmSessionableBindAttendee"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmSessionableBindAttendee
	err := bind.DeleteAll(c.(request.Request))
	return err
}
