package guardian

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmGuardian struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Relationship string `json:"Relationship" bson:"Relationship"`

	BrandId string `json:"brandId" bson:"brandId"`

	Name     string `json:"name" bson:"name"`
	Nickname string `json:"nickname" bson:"nickname"`
	Icon     string `json:"icon" bson:"icon"`
	Dob      int64  `json:"dob" bson:"dob"`
	Gender   int64  `json:"gender" bson:"gender"`
	RegDate  int64  `json:"reg_date" bson:"reg_date"`
	Contact  string `json:"contact" bson:"contact"`
	WeChat   string `json:"wechat" bson:"wechat"`

	//Person person.BmPerson `json:"Person" jsonapi:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmGuardian) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmGuardian) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmGuardian) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmGuardian) QueryId() string {
	return bd.Id
}

func (bd *BmGuardian) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmGuardian) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmGuardian) SetConnect(tag string, v interface{}) interface{} {
	//switch tag {
	//case "Person":
	//	bd.Person = v.(person.BmPerson)
	//}
	return bd
}

func (bd BmGuardian) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmGuardian) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmGuardian) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmGuardian) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

//func (bd *BmGuardian) ReSetPerson() error {
//
//	var err error
//
//	eq1 := request.Eqcond{}
//	eq1.Ky = "guardianId"
//	eq1.Vy = bd.Id
//	req1 := request.Request{}
//	req1.Res = "BMGuardianProp"
//	var condi1 []interface{}
//	condi1 = append(condi1, eq1)
//	c1 := req1.SetConnect("conditions", condi1)
//	guardianProp := BMGuardianProp{}
//	err = guardianProp.FindOne(c1.(request.Request))
//	if err != nil {
//		return err
//	}
//
//	eq := request.Eqcond{}
//	eq.Ky = "_id"
//	eq.Vy = bson.ObjectIdHex(guardianProp.PersonId)
//	req := request.Request{}
//	req.Res = "BmPerson"
//	var condi []interface{}
//	condi = append(condi, eq)
//	c := req.SetConnect("conditions", condi)
//	person := person.BmPerson{}
//	err = person.FindOne(c.(request.Request))
//	if err != nil {
//		return err
//	}
//	bd.Person = person
//
//	return nil
//}
