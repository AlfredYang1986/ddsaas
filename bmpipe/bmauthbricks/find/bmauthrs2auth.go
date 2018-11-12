package authfind

import (
	"fmt"
	//"github.com/alfredyang1986/ddsaas/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMAuthRS2AuthBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAuthRS2AuthBrick) Exec() error {
	prop := b.bk.Pr.(auth.BmAuthProp)
	reval, err := findAuth(prop)
	phone, err := findPhone(prop)
	wechat, err := findWechat(prop)
	//profile, err := findProfile(prop)
	reval.Phone = phone
	reval.Wechat = wechat
	//reval.Profile = profile
	b.bk.Pr = reval
	return err
}

func (b *BMAuthRS2AuthBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BmAuthProp)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAuthRS2AuthBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAuthRS2AuthBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAuthRS2AuthBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BmAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAuthRS2AuthBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmAuth = b.BrickInstance().Pr.(auth.BmAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

/*------------------------------------------------
 * brick inner function
 *------------------------------------------------*/

func findPhone(prop auth.BmAuthProp) (auth.BmPhone, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Phone_id)
	req := request.Request{}
	req.Res = "BmPhone"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BmPhone{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findWechat(prop auth.BmAuthProp) (auth.BmWeChat, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Wechat_id)
	req := request.Request{}
	req.Res = "BmWeChat"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BmWeChat{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findAuth(prop auth.BmAuthProp) (auth.BmAuth, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Auth_id)
	req := request.Request{}
	req.Res = "BmAuth"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BmAuth{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}
