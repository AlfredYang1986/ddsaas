package sessionable

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmSessionableBindTeacher struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	SessionableId string `json:"sessionableId" bson:"sessionableId"`
	TeacherId     string `json:"teacherId" bson:"teacherId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmSessionableBindTeacher) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmSessionableBindTeacher) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmSessionableBindTeacher) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmSessionableBindTeacher) QueryId() string {
	return bd.Id
}

func (bd *BmSessionableBindTeacher) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmSessionableBindTeacher) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmSessionableBindTeacher) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmSessionableBindTeacher) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmSessionableBindTeacher) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmSessionableBindTeacher) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmSessionableBindTeacher) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmSessionableBindTeacher) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmSessionableBindTeacher) DeleteAll(req request.Request) error {
	return bmmodel.DeleteAll(req, bd)
}

func (bd *BmSessionableBindTeacher) Clear() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "sessionableId"
	eq1.Vy = bd.SessionableId
	req := request.Request{}
	req.Res = "BmSessionableBindTeacher"
	var condi []interface{}
	condi = append(condi, eq1)
	c := req.SetConnect("conditions", condi)
	var bind BmSessionableBindTeacher
	err := bind.DeleteAll(c.(request.Request))
	return err
}
