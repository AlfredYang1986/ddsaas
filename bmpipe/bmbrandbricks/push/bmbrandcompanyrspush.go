package brandpush

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"github.com/alfredyang1986/ddsaas/bmmodel/profile"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMBrandCompanyRSPush struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMBrandCompanyRSPush) Exec() error {
	var tmp brand.BMBrand = b.bk.Pr.(brand.BMBrand)

	company, err := findCompany(tmp)
	if err != nil {
		b.bk.Err = -6
		return err
	}
	tmp.Company = company
	eq := request.EQCond{}
	eq.Ky = "brand_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BMBrandCompanyRS"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var qr brand.BMBrandCompanyRS
	err = qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.BrandId = tmp.Id
		qr.CompanyId = company.Id
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	b.bk.Pr = tmp
	return nil
}

func (b *BMBrandCompanyRSPush) Prepare(pr interface{}) error {
	req := pr.(brand.BMBrand)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMBrandCompanyRSPush) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMBrandCompanyRSPush) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMBrandCompanyRSPush) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(brand.BMBrand)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMBrandCompanyRSPush) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval brand.BMBrand = b.BrickInstance().Pr.(brand.BMBrand)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func findCompany(b brand.BMBrand) (profile.BMCompany, error) {
	eq := request.EQCond{}
	eq.Ky = "name"
	eq.Vy = b.Company.Name
	req := request.Request{}
	req.Res = "BMCompany"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := profile.BMCompany{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}

