package attendeeupdate

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"io"
	"net/http"
)

type BmAttendeeUpdateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmAttendeeUpdateBrick) Exec() error {

	attendeeReq := b.bk.Req
	attendee := attendee.BmAttendee{}
	attendee.UpdateBMObject(*attendeeReq)

	//guardianReq := b.bk.Req
	//err, attendeeGuardianRSes := attendee.GetAttendeeGuardianRSes()
	//guardianReq.Res = "BmGuardian"
	//guardians := []guardian.BmGuardian{}
	//for _,v := range attendeeGuardianRSes {
	//	errtmp, g := v.GetGuardian()
	//	if errtmp != nil {
	//		return errtmp
	//	}
	//	g.UpdateBMObject(*guardianReq)
	//	guardians = append(guardians, g)
	//}
	//attendee.Guardians = guardians

	b.BrickInstance().Pr = attendee
	return nil
}

func (b *BmAttendeeUpdateBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmAttendeeUpdateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmAttendeeUpdateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmAttendeeUpdateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BmAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmAttendeeUpdateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(attendee.BmAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
