package brandpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmBrandBindProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBrandBindProp) Exec() error {
	tmp := b.bk.Pr.(brand.BmBrand)

	cat := tmp.Cate
	bind := brand.BmBindBrandCategory{}
	bind.CategoryId = cat.Id
	bind.BrandId = tmp.Id
	bind.Id_ = bson.NewObjectId()
	bind.Id = bind.Id_.Hex()
	err := bind.InsertBMObject()

	for _, item := range tmp.Honors {
		ist := brand.BmBindBrandHonor{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.HonorId = item.Id
		ist.BrandId = tmp.Id
		ist.InsertBMObject()
	}

	for _, item := range tmp.Certifications {
		ist := brand.BmBindBrandCertific{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.CertificationId = item.Id
		ist.BrandId = tmp.Id
		ist.InsertBMObject()
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmBrandBindProp) Prepare(pr interface{}) error {
	req := pr.(brand.BmBrand)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBrandBindProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBrandBindProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBrandBindProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(brand.BmBrand)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBrandBindProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval brand.BmBrand = b.BrickInstance().Pr.(brand.BmBrand)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

