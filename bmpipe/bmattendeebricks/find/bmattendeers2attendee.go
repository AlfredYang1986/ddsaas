package attendeefind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"io"
	"net/http"
)

type BMAttendeeRS2Attendee struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeeRS2Attendee) Exec() error {
	tmp := b.bk.Pr.(attendee.BmAttendee)
	var err error

	eq2 := request.Eqcond{}
	var eq2arr []request.Eqcond
	eq2.Ky = "attendeeId"
	eq2.Vy = tmp.Id
	eq2.Ct = "BMAttendeeGuardianRS"
	req2 := request.Request{}
	req2.Res = "BMAttendeeGuardianRS"
	req2.Eqcond = append(eq2arr, eq2)
	var condi2 []interface{}
	condi2 = append(condi2, eq2)
	c2 := req2.SetConnect("Eqcond", condi2)
	var agrsarr attendee.BMAttendeeGuardianRSeS
	err = agrsarr.FindMulti(c2.(request.Request))
	if err != nil {
		return err
	}

	var guardians []guardian.BmGuardian
	for _, agr := range agrsarr.AgRsArr {
		err, g := agr.GetGuardian()
		if err != nil {
			return err
		}
		guardians = append(guardians, g)
	}
	tmp.Guardians = guardians

	b.bk.Pr = tmp
	return err
}

func (b *BMAttendeeRS2Attendee) Prepare(pr interface{}) error {
	req := pr.(attendee.BmAttendee)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAttendeeRS2Attendee) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeeRS2Attendee) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeeRS2Attendee) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BmAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeeRS2Attendee) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(attendee.BmAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

