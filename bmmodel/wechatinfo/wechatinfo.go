package wechatinfo

import (
	"encoding/json"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
)

type BmWeChatInfo struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	AppId      string `json:"AppId" bson:"AppId"`
	Secret     string `json:"Secret" bson:"Secret"`
	OpenId     string `json:"OpenId" bson:"OpenId"`
	SessionKey string `json:"SessionKey" bson:"SessionKey"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmWeChatInfo) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmWeChatInfo) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmWeChatInfo) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmWeChatInfo) QueryId() string {
	return bd.Id
}

func (bd *BmWeChatInfo) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmWeChatInfo) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmWeChatInfo) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmWeChatInfo) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmWeChatInfo) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmWeChatInfo) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *BmWeChatInfo) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmWeChatInfo) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmWeChatInfo) GetWeChatInfo(req request.Request) error {

	if len(req.Eqcond)<1 {
		panic("no code")
	}
	var code string
	for _,eq := range req.Eqcond {
		if eq.Ky == "code" {
			code = eq.Vy.(string)
		}
		if eq.Ky == "brand" {
			brand := eq.Vy.(string)
			bd.getBrandAppId(brand)
		}
	}

	originUrl := "https://api.weixin.qq.com/sns/jscode2session?appid="
	url := strings.Join([]string{originUrl, bd.AppId, "&secret=", bd.Secret, "&js_code=", code, "&grant_type=authorization_code"}, "")
	resp, err := http.Get(url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(body, &m)

	openid := m["openid"]
	session_key := m["session_key"]

	bd.OpenId = openid
	bd.SessionKey = session_key

	return err
}

func (bd *BmWeChatInfo) getBrandAppId(brand string) error {

	switch brand {
	case "dongda":
		bd.AppId = "wx6129e48a548c52b8"
		bd.Secret = "b250e875e51a931e2ae3a49ff450bc3c"
	case "pacee":
		bd.AppId = "wx79138b2ee5288cc2"
		bd.Secret = "c2637375412cfa97c9e127b4cde30c5c"
	}
	return nil
}
