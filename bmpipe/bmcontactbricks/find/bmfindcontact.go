package contactfind

import (
	"github.com/alfredyang1986/blackmirror-modules/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmmodel/contact"
	"github.com/alfredyang1986/blackmirror-modules/bmmodel/request"
	"github.com/alfredyang1986/blackmirror-modules/bmerror"
	"github.com/alfredyang1986/blackmirror-modules/bmpipe"
	"github.com/alfredyang1986/blackmirror-modules/bmrouter"
	"github.com/alfredyang1986/blackmirror-modules/jsonapi"
	"net/http"
	"io"
	"fmt"
)

type BMContactFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMContactFindBrick) Exec() error {
	var tmp contact.Contact
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BMContactFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	fmt.Println(req)
	b.BrickInstance().Req = &req
	//b.bk.Pr = req
	return nil
}

func (b *BMContactFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMContactFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMContactFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(contact.Contact)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMContactFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval contact.Contact = b.BrickInstance().Pr.(contact.Contact)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
