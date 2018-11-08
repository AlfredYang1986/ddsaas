package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/reward"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMBrand struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title      string   `json:"title" bson:"title"`
	Subtitle   string   `json:"subtitle" bson:"subtitle"`
	BrandTags  []string `json:"brand_tags" bson:"brand_tags"`
	Found      int64    `json:"found"`
	FoundStory string   `json:"FoundStory" bson:"FoundStory"`

	Rewards  []reward.BMReward   `json:"rewards" jsonapi:"relationships"`
	//Students []student.BMStudent `json:"students" jsonapi:"relationships"`
	//Attendees []attendee.BmAttendee `json:"attendees" jsonapi:"relationships"`
	//Teachers []teacher.BmTeacher    `json:"teachers" jsonapi:"relationships"`
	//Sales    []sales.BMSales        `json:"sales" jsonapi:"relationships"`

	//Yard []yard.BMYard `json:"yard" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMBrand) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMBrand) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMBrand) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMBrand) QueryId() string {
	return bd.Id
}

func (bd *BMBrand) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMBrand) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMBrand) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "rewards":
		var rst []reward.BMReward
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(reward.BMReward))
		}
		bd.Rewards = rst
	//case "students":
	//	var rst []student.BMStudent
	//	for _, item := range v.([]interface{}) {
	//		rst = append(rst, item.(student.BMStudent))
	//	}
	//	bd.Students = rst
	//case "attendees":
	//	var rst []attendee.BmAttendee
	//	for _, item := range v.([]interface{}) {
	//		rst = append(rst, item.(attendee.BmAttendee))
	//	}
	//	bd.Attendees = rst
	//case "teachers":
	//	var rst []teacher.BmTeacher
	//	for _, item := range v.([]interface{}) {
	//		rst = append(rst, item.(teacher.BmTeacher))
	//	}
	//	bd.Teachers = rst
	//case "sales":
	//	var rst []sales.BMSales
	//	for _, item := range v.([]interface{}) {
	//		rst = append(rst, item.(sales.BMSales))
	//	}
	//	bd.Sales = rst
	//case "yard":
	//	var rst []yard.BMYard
	//	for _, item := range v.([]interface{}) {
	//		rst = append(rst, item.(yard.BMYard))
	//	}
	//	bd.Yard = rst
	}
	return bd
}

func (bd BMBrand) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMBrand) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMBrand) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMBrand) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd BMBrand) IsBrandRegistered() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMBrand")
	n, err := c.Find(bson.M{"title": bd.Title}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMBrand) Valid() bool {
	return bd.Title != ""
}
