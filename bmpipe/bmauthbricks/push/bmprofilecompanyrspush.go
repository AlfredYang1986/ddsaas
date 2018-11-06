package authpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"io"
	"net/http"
)

type BMProfileCompanyRSPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMProfileCompanyRSPushBrick) Exec() error {
	//var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)

	//p := tmp.Profile
	//company := p.Company
	//eq := request.EQCond{}
	//eq.Ky = "profile_id"
	//eq.Vy = p.Id
	//req := request.Request{}
	//req.Res = "BMProfileCompanyRS"
	//var condi []interface{}
	//condi = append(condi, eq)
	//c := req.SetConnect("conditions", condi)
	//fmt.Println(c)
	//
	//var qr profile.BMProfileCompanyRS
	//err := qr.FindOne(c.(request.Request))
	//if err != nil && err.Error() == "not found" {
	//	qr.Id_ = bson.NewObjectId()
	//	qr.Id = qr.Id_.Hex()
	//	qr.ProfileId = p.Id
	//	qr.CompanyId = company.Id
	//	qr.InsertBMObject()
	//}
	//fmt.Println(qr)
	return nil
}

func (b *BMProfileCompanyRSPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMProfileCompanyRSPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMProfileCompanyRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMProfileCompanyRSPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMProfileCompanyRSPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
