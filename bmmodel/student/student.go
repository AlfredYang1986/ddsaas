package student

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMStudent struct {
	Id        			string            	`json:"id"`
	Id_       			bson.ObjectId     	`bson:"_id"`

	Name      			string            	`json:"name" bson:"name"`
	NickName  			string           	`json:"nickname" bson:"nickname"`
	Birthday  			string           	`json:"birthday" bson:"birthday"`
	Age  				int           		`json:"age" bson:"age"`
	Sex  				string           	`json:"sex" bson:"sex"`
	School  			string           	`json:"school" bson:"school"`
	Photo  				string           	`json:"photo" bson:"photo"`

	Guardians			[]BMGuardian		`json:"guardians" jsonapi:"relationships"`
	Contacts			[]BMContacter		`json:"contacts" jsonapi:"relationships"`

	//Found     			time.Time       	`json:"found"`
	//Patriarch 			map[string]interface{} 	`json:"patriarch" bson:"patriarch"`			//del

	//ContinuedCourses 	[]course.BMCourse 	`json:"continuedcourses" jsonapi:"relationships"`	//設計邏輯還未明確
	//CompletedCourses 	[]course.BMCourse 	`json:"completedcourses" jsonapi:"relationships"`	//設計邏輯還未明確

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMStudent) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMStudent) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMStudent) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMStudent) QueryId() string {
	return bd.Id
}

func (bd *BMStudent) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMStudent) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMStudent) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "guardians":
		var guardians []BMGuardian
		for _, item := range v.([]interface{}) {
			guardians = append(guardians, item.(BMGuardian))
		}
		bd.Guardians = guardians
	case "orders":
		var contacts []BMContacter
		for _, item := range v.([]interface{}) {
			contacts = append(contacts, item.(BMContacter))
		}
		bd.Contacts = contacts
	}
	return bd
}

func (bd BMStudent) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMStudent) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMStudent) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMStudent) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
