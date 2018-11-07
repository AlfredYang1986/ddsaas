package personupdate

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"io"
	"net/http"
)

type BmPersonUpdate struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmPersonUpdate) Exec() error {
	tmp := person.BmPerson{}
	tmp.UpdateBMObject(*b.bk.Req)
	b.BrickInstance().Pr = tmp
	return nil
}

func (b *BmPersonUpdate) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Req = &req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmPersonUpdate) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmPersonUpdate) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmPersonUpdate) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(person.BmPerson)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmPersonUpdate) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval person.BmPerson = b.BrickInstance().Pr.(person.BmPerson)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
