package contactpush

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/contact"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMContactRSPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMContactRSPushBrick) Exec() error {
	var tmp = b.bk.Pr.(contact.Contact)
	eq := request.EQCond{}
	eq.Ky = "contact_id"
	eq.Vy = tmp.Id
	var tmpOrderIds []string
	for _, v := range tmp.Orders {
		tmpOrderIds = append(tmpOrderIds, v.Id)
	}
	req := request.Request{}
	req.Res = "BMContactProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var qr contact.BMContactProp
	err := qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.ContactId = tmp.Id
		qr.OrderIds = tmpOrderIds
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	return nil
}

func (b *BMContactRSPushBrick) Prepare(pr interface{}) error {
	req := pr.(contact.Contact)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMContactRSPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMContactRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMContactRSPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(contact.Contact)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMContactRSPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval = b.BrickInstance().Pr.(contact.Contact)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
