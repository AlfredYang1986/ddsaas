package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmCourseUnits struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseUnits []BmCourseUnit `json:"CourseUnits" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnits) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnits) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnits) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnits) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnits) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCourseUnits) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnits) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnits) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnits) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.CourseUnits)
	for i, r := range bd.CourseUnits {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.CourseUnits[i] = r
	}
	return err
}
