package attendeepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMAttendeePushPersonRS struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeePushPersonRS) Exec() error {
	var tmp attendee.BMAttendee = b.bk.Pr.(attendee.BMAttendee)

	person := tmp.Person
	var ap attendee.BMAttendeeProp
	ap.Id_ = bson.NewObjectId()
	ap.Id = ap.Id_.Hex()
	ap.AttendeeId = tmp.Id
	ap.PersonId = person.Id
	ap.InsertBMObject()

	b.bk.Pr = tmp
	return nil
}

func (b *BMAttendeePushPersonRS) Prepare(pr interface{}) error {
	req := pr.(attendee.BMAttendee)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAttendeePushPersonRS) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeePushPersonRS) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeePushPersonRS) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BMAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeePushPersonRS) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval attendee.BMAttendee = b.BrickInstance().Pr.(attendee.BMAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
