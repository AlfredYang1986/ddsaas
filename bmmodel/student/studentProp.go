package student

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMStudentProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	StudentID string `json:"student_id" bson:"student_id"`

	GuardianIds []interface{} `json:"guardianids" bson:"guardianids"`
	ContactIds  []interface{} `json:"contactids" bson:"contactids"`
	//ContinuedCoursesIds   	[]string 		`json:"continued_courses_ids" bson:"continued_courses_ids"`
	//CompletedCoursesIds   	[]string 		`json:"completed_courses_ids" bson:"completed_courses_ids"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMStudentProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMStudentProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMStudentProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMStudentProp) QueryId() string {
	return bd.Id
}

func (bd *BMStudentProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMStudentProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMStudentProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMStudentProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMStudentProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMStudentProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMStudentProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
