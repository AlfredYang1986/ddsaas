package attendeepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BMAttendeePushPerson struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeePushPerson) Exec() error {
	var tmp attendee.BMAttendee = b.bk.Pr.(attendee.BMAttendee)
	ap := tmp.Person
	ap.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BMAttendeePushPerson) Prepare(pr interface{}) error {
	req := pr.(attendee.BMAttendee)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAttendeePushPerson) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeePushPerson) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeePushPerson) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BMAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeePushPerson) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval attendee.BMAttendee = b.BrickInstance().Pr.(attendee.BMAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

