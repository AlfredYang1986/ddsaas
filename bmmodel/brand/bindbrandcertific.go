package brand

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmBindBrandCertific struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	BrandId         string `json:"brandId" bson:"brandId"`
	CertificationId string `json:"certificationId" bson:"certificationId"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmBindBrandCertific) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmBindBrandCertific) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmBindBrandCertific) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmBindBrandCertific) QueryId() string {
	return bd.Id
}

func (bd *BmBindBrandCertific) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmBindBrandCertific) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmBindBrandCertific) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmBindBrandCertific) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmBindBrandCertific) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmBindBrandCertific) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmBindBrandCertific) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmBindBrandCertific) CheckExist() error {

	eq1 := request.Eqcond{}
	eq1.Ky = "certificationId"
	eq1.Vy = bd.CertificationId
	eq2 := request.Eqcond{}
	eq2.Ky = "brandId"
	eq2.Vy = bd.BrandId
	req := request.Request{}
	req.Res = "BmBindBrandCertific"
	var condi []interface{}
	condi = append(condi, eq1)
	condi = append(condi, eq2)
	c := req.SetConnect("conditions", condi)
	var bind BmBindBrandCertific
	err := bind.FindOne(c.(request.Request))
	if bind.Id != "" {
		bd.SetId(bind.Id)
		bd.SetObjectId(bind.Id_)
	}
	return err
}
