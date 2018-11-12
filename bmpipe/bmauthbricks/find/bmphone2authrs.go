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
	"io"
	"net/http"
)

type BMPhone2AuthRSBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMPhone2AuthRSBrick) Exec() error {
	var tmp auth.BmPhone = b.bk.Pr.(auth.BmPhone)
	eq := request.Eqcond{}
	eq.Ky = "phone_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BmAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var reval auth.BmAuthProp
	err := reval.FindOne(c.(request.Request))
	b.bk.Pr = reval
	return err
}

func (b *BMPhone2AuthRSBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BmPhone)
	b.BrickInstance().Pr = req
	//b.bk.Pr = req
	return nil
}

func (b *BMPhone2AuthRSBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMPhone2AuthRSBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMPhone2AuthRSBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BmAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMPhone2AuthRSBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmAuth = b.BrickInstance().Pr.(auth.BmAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
