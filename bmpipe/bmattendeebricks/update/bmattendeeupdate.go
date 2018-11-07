package attendeeupdate

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"io"
	"net/http"
)

type BmAttendeeUpdate struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmAttendeeUpdate) Exec() error {

	attendeeReq := b.bk.Req
	personReq := b.bk.Req

	attendee := attendee.BmAttendee{}
	person := person.BmPerson{}

	attendee.UpdateBMObject(*attendeeReq)

	personReq.Res = "BmPerson"
	person.UpdateBMObject(*personReq)
	attendee.Person = person

	b.BrickInstance().Pr = attendee
	return nil
}

func (b *BmAttendeeUpdate) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmAttendeeUpdate) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmAttendeeUpdate) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmAttendeeUpdate) ResultTo(w io.Writer) error {
	//pr := b.BrickInstance().Pr
	//tmp := pr.(attendee.BmAttendee)
	//err := jsonapi.ToJsonAPI(&tmp, w)

	tmp := b.BrickInstance().Req
	tmp.Res = "BmPerson"
	err := jsonapi.ToJsonAPI(tmp, w)

	return err
}

func (b *BmAttendeeUpdate) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(attendee.BmAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}