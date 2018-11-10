package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"github.com/alfredyang1986/ddsaas/bmmodel/certification"
	"github.com/alfredyang1986/ddsaas/bmmodel/honor"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BmBrand struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Title      string `json:"title" bson:"title"`
	Subtitle   string `json:"subtitle" bson:"subtitle"`
	Found      int64  `json:"found"`
	FoundStory string `json:"foundStory" bson:"foundStory"`

	//TODO:20181109新增的
	Cate      category.BmCategory `json:"Cate" jsonapi:"relationships"` //类别
	Logo      string              `json:"logo" bson:"logo"`             //品牌logo
	Slogan    string              `json:"slogan" bson:"slogan"`         //一句话介绍
	BrandTags []interface{}       `json:"brand_tags" bson:"brand_tags"` //HightLight[与众不同],3-5条,一条5个字
	EduIdea   string              `json:"edu_idea" bson:"edu_idea"`     //教育理念
	AboutUs   string              `json:"about_us" bson:"about_us"`     //团队
	//TODO:Honors和Certifications合并成TagImgs,添加category做区分.
	Honors         []honor.BmHonor                 `json:"Honors" jsonapi:"relationships"`
	Certifications []certification.BmCertification `json:"Certifications" jsonapi:"relationships"`

	//Students []student.BMStudent `json:"students" jsonapi:"relationships"`
	//Attendees []attendee.BmAttendee `json:"attendees" jsonapi:"relationships"`
	//Teachers []teacher.BmTeacher    `json:"teachers" jsonapi:"relationships"`
	//Sales    []sales.BMSales        `json:"sales" jsonapi:"relationships"`

	//Yard []yard.BMYard `json:"yard" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBrand) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBrand) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBrand) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBrand) QueryId() string {
	return bd.Id
}

func (bd *BmBrand) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBrand) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBrand) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Cate":
		bd.Cate = v.(category.BmCategory)
	case "Honors":
		var rst []honor.BmHonor
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(honor.BmHonor))
		}
		bd.Honors = rst
	case "Certifications":
		var rst []certification.BmCertification
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(certification.BmCertification))
		}
		bd.Certifications = rst
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

func (bd BmBrand) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBrand) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBrand) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBrand) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd BmBrand) IsBrandRegistered() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()
	c := session.DB("test").C("BmBrand")

	n, err := c.Find(bson.M{"title": bd.Title}).Count()
	if err != nil {
		panic(err)
	}
	return n > 0
}

func (bd BmBrand) Valid() bool {
	return bd.Title != ""
}

func (bd *BmBrand) ReSetProp() error {

	bd.reSetCate()
	bd.reSetHonor()
	bd.reSetCertification()

	return nil
}

func (bd *BmBrand) reSetCate() error {

	eq := request.Eqcond{}
	eq.Ky = "brandId"
	eq.Vy = bd.Id
	req := request.Request{}
	req.Res = "BmBindBrandCategory"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	reval := BmBindBrandCategory{}
	err := reval.FindOne(c.(request.Request))

	eq0 := request.Eqcond{}
	eq0.Ky = "id"
	eq0.Vy = reval.CategoryId
	req0 := request.Request{}
	req0.Res = "BmCategory"
	var condi0 []interface{}
	condi0 = append(condi0, eq0)
	c0 := req0.SetConnect("conditions", condi0)

	result := category.BmCategory{}
	err = result.FindOne(c0.(request.Request))
	bd.Cate = result

	return err
}
func (bd *BmBrand) reSetHonor() error {

	req := request.Request{}
	req.Res = "BmBindBrandHonor"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "brandId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindBrandHonor
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.HonorId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []honor.BmHonor
	err = bmmodel.FindMutilWithBson("BmHonor", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		resultArr[i] = ir
	}
	bd.Honors = resultArr

	return nil
}
func (bd *BmBrand) reSetCertification() error {

	req := request.Request{}
	req.Res = "BmBindBrandCertific"
	var condi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "brandId"
	eq.Vy = bd.Id
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	var reval []BmBindBrandCertific
	err := bmmodel.FindMutil(c.(request.Request), &reval)
	if err != nil {
		return err
	}

	var condi0 []bson.ObjectId
	for _, item := range reval {
		condi0 = append(condi0, bson.ObjectIdHex(item.CertificationId))
	}

	tt := make(map[string]interface{})
	tt["$in"] = condi0
	or_condi := bson.M{"_id": tt}

	var resultArr []certification.BmCertification
	err = bmmodel.FindMutilWithBson("BmCertification", or_condi, &resultArr)

	for i, ir := range resultArr {
		ir.ResetIdWithId_()
		resultArr[i] = ir
	}
	bd.Certifications = resultArr

	return nil
}
