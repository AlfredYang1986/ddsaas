package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnit struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Status      float64                   `json:"status" bson:"status"`
	StartDate   float64                   `json:"start_date" bson:"start_date"`
	EndDate     float64                   `json:"end_date" bson:"end_date"`
	Sessionable sessionable.BmSessionable `json:"Sessionable" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnit) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnit) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnit) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnit) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnit) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Sessionable":
		bd.Sessionable = v.(sessionable.BmSessionable)
	}
	return bd
}

func (bd BmCourseUnit) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnit) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnit) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmCourseUnit) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnit) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnit) ReSetProp() error {
	bd.reSetSessionabel()
	return nil
}

func (bd *BmCourseUnit) reSetSessionabel() error {

	eq := request.Eqcond{}
	eq.Ky = "courseUnitId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmCourseUnitBindSessionable"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmCourseUnitBindSessionable{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.SessionableId
	req0 := request.Request{}
	req0.Res = "BmSessionable"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := sessionable.BmSessionable{}
	err = result.FindOne(c0.(request.Request))
	result.ReSetProp()
	bd.Sessionable = result

	return err
}
