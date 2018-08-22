package course

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMCourse struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseName   string        `json:"coursename" bson:"course_name"`
	Tags         []interface{} `json:"tags" bson:"tags"`
	Description  string        `json:"description" bson:"description"`
	CourseType   string        `json:"coursetype" bson:"course_type"`
	MinAge       int           `json:"minage" bson:"min_age"`
	MaxAge       int           `json:"maxage" bson:"max_age"`
	Level        string        `json:"level" bson:"level"`
	PeriodCount  int           `json:"periodcount" bson:"period_count"`
	Duration     int           `json:"duration" bson:"duration"`
	Mode         []interface{} `json:"mode" bson:"mode"`
	Aims         string        `json:"aims" bson:"aims"`
	Outline      string        `json:"outline" bson:"outline"`
	Introduction string        `json:"introduction" bson:"introduction"`
	CoursePhoto  []interface{} `json:"coursephoto" bson:"course_photo"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMCourse) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMCourse) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMCourse) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMCourse) QueryId() string {
	return bd.Id
}

func (bd *BMCourse) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMCourse) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMCourse) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMCourse) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMCourse) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMCourse) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMCourse) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * course verified interface
 *------------------------------------------------*/

func (bd BMCourse) IsRegistered() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMCourse")
	n, err := c.Find(bson.M{"course_name": bd.CourseName}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMCourse) Valid() bool {
	return bd.CourseName != ""
}
