package activity

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMActivity struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	//common info
	Name        		string        		`json:"name" bson:"name"`
	Type        		string       	 	`json:"type" bson:"type"`													//course or experience_class
	MinAge      		int           		`json:"minage" bson:"min_age"`
	MaxAge      		int           		`json:"maxage" bson:"max_age"`
	Duration    		int           		`json:"duration" bson:"duration"`											//持续时间【课程时长/体验时长】
	Introduction 		string        		`json:"introduction" bson:"introduction"`									//介绍 课程/活动
	Tags        		[]interface{} 		`json:"tags" bson:"tags"`
	Photos  			[]interface{} 		`json:"photos" bson:"photos"`
	Address				string				`json:"address" bson:"address"`
	Size       			int    				`json:"size" bson:"size"`													//名额
	Deadline      		string 				`json:"deadline" bson:"deadline"`											//报名截至日期
	StartDate      		string 				`json:"startdate" bson:"start_date"`
	EndDate        		string 				`json:"enddate" bson:"end_date"`
	StartTime      		string 				`json:"starttime" bson:"start_time"`
	EndTime      		string 				`json:"endtime" bson:"end_time"`

	Brand				brand.BMBrand		`json:"brand" jsonapi:"relationships"`

	//experience_class info
	ProcessDesign 		string        		`json:"processdesign" bson:"process_design"`
	BabyRewards			[]interface{}		`json:"babyrewards" bson:"baby_rewards"`
	BusinessProvides	[]interface{}		`json:"businessprovides" bson:"business_provides"`
	ParticipantProvides	[]interface{}		`json:"participantprovides" bson:"participant_provides"`
	Notice 				string        		`json:"notice" bson:"notice"`												//要求/须知
	Kostenindicatie 	string        		`json:"kostenindicatie" bson:"kostenindicatie"`								//费用说明

	//course info
	CourseType  		string        		`json:"coursetype" bson:"course_type"`
	Description 		string        		`json:"description" bson:"description"`
	Level       		string        		`json:"level" bson:"level"`
	PeriodCount 		int           		`json:"periodcount" bson:"period_count"`
	Mode         		[]interface{} 		`json:"mode" bson:"mode"`
	Aims         		string        		`json:"aims" bson:"aims"`
	Outline      		string        		`json:"outline" bson:"outline"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMActivity) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMActivity) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMActivity) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMActivity) QueryId() string {
	return bd.Id
}

func (bd *BMActivity) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMActivity) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMActivity) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "brand":
		bd.Brand = v.(brand.BMBrand)
	}
	return bd
}

func (bd BMActivity) QueryConnect(tag string) interface{} {
	switch tag {
	case "brand":
		return bd.Brand
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMActivity) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMActivity) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMActivity) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * course verified interface
 *------------------------------------------------*/

func (bd BMActivity) IsRegistered() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMActivity")
	n, err := c.Find(bson.M{"name": bd.Name}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMActivity) Valid() bool {
	return bd.Name != ""
}
