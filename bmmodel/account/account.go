package account

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmsecurity"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
)

type BMAccount struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Account    string `json:"account" bson:"account"`
	SecretWord string `json:"secretword" bson:"secretword"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAccount) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAccount) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAccount) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAccount) QueryId() string {
	return bd.Id
}

func (bd *BMAccount) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAccount) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAccount) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAccount) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAccount) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAccount) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAccount) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMAccount) DecodeByCompanyDate(company string, date string) error {

	var bmRsaKey auth.BMRsaKey = auth.BMRsaKey{
		Company:company,
		Date:date,
	}

	privateKey, err := bmRsaKey.GetPrivateKey()
	if err != nil {
		return err
	}

	secretWord := bd.SecretWord
	secretByte, err := base64.StdEncoding.DecodeString(secretWord)
	if err != nil {
		return err
	}

	originByte, err := bmsecurity.PhRsaDecrypt(privateKey, secretByte)
	if err != nil {
		return err
	}

	bd.SecretWord = string(originByte)

	return nil
}

func (bd *BMAccount) Secret2MD5() {

	secretWord := bd.SecretWord

	h := md5.New()
	io.WriteString(h, secretWord)

	secretWordMd5 := fmt.Sprintf("%x", h.Sum(nil))
	bd.SecretWord = secretWordMd5

}

func (bd BMAccount) IsAccountRegisted() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMAccount")
	n, err := c.Find(bson.M{"account": bd.Account}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMAccount) Valid() bool {
	return bd.Account != ""
}
