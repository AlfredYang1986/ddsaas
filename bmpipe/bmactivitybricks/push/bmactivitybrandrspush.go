package activitypush

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/activity"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMActivityBrandRSPush struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMActivityBrandRSPush) Exec() error {
	var tmp activity.BMActivity = b.bk.Pr.(activity.BMActivity)

	brand_tmp, err := findBrand(tmp)
	if err != nil {
		b.bk.Err = -7
		return err
	}
	tmp.Brand = brand_tmp

	eq := request.EQCond{}
	eq.Ky = "activity_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BMActivityBrandRS"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var qr activity.BMActivityBrandRS
	err = qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.ActivityId = tmp.Id
		qr.BrandId = brand_tmp.Id
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	b.bk.Pr = tmp
	return nil
}

func (b *BMActivityBrandRSPush) Prepare(pr interface{}) error {
	req := pr.(activity.BMActivity)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMActivityBrandRSPush) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMActivityBrandRSPush) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMActivityBrandRSPush) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(activity.BMActivity)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMActivityBrandRSPush) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval activity.BMActivity = b.BrickInstance().Pr.(activity.BMActivity)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func findBrand(a activity.BMActivity) (brand.BMBrand, error) {
	eq := request.EQCond{}
	eq.Ky = "name"
	//eq.Vy = a.Brand.Name
	req := request.Request{}
	req.Res = "BMBrand"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := brand.BMBrand{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}

