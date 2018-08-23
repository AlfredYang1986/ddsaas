package class

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMClass struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	ClassName	  	string 				`json:"classname" bson:"classname"`
	CourseName	  	string 				`json:"coursename" bson:"coursename"`
	Address	  		string 				`json:"address" bson:"address"`
	Size	  		int 				`json:"size" bson:"size"`
	Start	  		string 				`json:"start" bson:"start"`
	End	  			string				`json:"end" bson:"end"`

	Lessons			int					`json:"lessons" bson:"lessons"`		//课节数
	PeriodCount		int					`json:"periodcount" bson:"periodcount"`		//课时数

	Teachers		[]interface{}		`json:"teachers" bson:"teachers"`
	CourseDetail	interface{}			`json:"coursedetail" bson:"coursedetail"`
	Students		[]interface{}		`json:"students" bson:"students"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMClass) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMClass) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMClass) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMClass) QueryId() string {
	return bd.Id
}

func (bd *BMClass) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMClass) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMClass) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMClass) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMClass) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMClass) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMClass) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}