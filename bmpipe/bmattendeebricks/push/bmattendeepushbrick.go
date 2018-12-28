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
	"time"
)

type BMAttendeePushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeePushBrick) Exec() error {
	var tmp attendee.BmAttendee = b.bk.Pr.(attendee.BmAttendee)

	//TODO:暂时把cover逻辑写在InsertBMObject()中.
	//if tmp.IsAttendeeExist() {
	//	tmp.CoverBMObject()
	//} else {
	//	tmp.InsertBMObject()
	//}
	tmp.CreateTime = time.Now().UnixNano() / 1e6
	tmp.InsertBMObject()

	b.bk.Pr = tmp
	return nil
}

func (b *BMAttendeePushBrick) Prepare(pr interface{}) error {
	req := pr.(attendee.BmAttendee)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAttendeePushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeePushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BmAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval attendee.BmAttendee = b.BrickInstance().Pr.(attendee.BmAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

