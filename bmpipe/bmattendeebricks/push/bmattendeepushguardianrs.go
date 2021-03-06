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

type BMAttendeePushGuardianRS struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeePushGuardianRS) Exec() error {
	var tmp attendee.BmAttendee = b.bk.Pr.(attendee.BmAttendee)
	var err error

	guardians := tmp.Guardians
	applyees := tmp.Applyees

	agrs := attendee.BMAttendeeGuardianRS{}
	agrs.Id_ = bson.NewObjectId()
	agrs.Id = agrs.Id_.Hex()
	agrs.AttendeeId = tmp.Id
	agrs.Clear()

	//aba := attendee.BMAttendeeBindApplyee{}
	//aba.Id_ = bson.NewObjectId()
	//aba.Id = aba.Id_.Hex()
	//aba.AttendeeId = tmp.Id
	//aba.Clear()

	for _,g := range guardians{
		var ag attendee.BMAttendeeGuardianRS
		ag.Id_ = bson.NewObjectId()
		ag.Id = ag.Id_.Hex()
		ag.AttendeeId = tmp.Id
		ag.GuardianId = g.Id
		ag.InsertBMObject()
	}
	for _,ape := range applyees{
		var bind attendee.BMAttendeeBindApplyee
		bind.Id_ = bson.NewObjectId()
		bind.Id = bind.Id_.Hex()
		bind.AttendeeId = tmp.Id
		bind.ApplyeeId = ape.Id
		bind.CheckExist()
		err = bind.CheckExist()
		if err.Error() != "Already Exist!" {
			bind.InsertBMObject()
		}
	}
	b.bk.Pr = tmp
	return nil
}

func (b *BMAttendeePushGuardianRS) Prepare(pr interface{}) error {
	req := pr.(attendee.BmAttendee)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAttendeePushGuardianRS) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeePushGuardianRS) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeePushGuardianRS) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BmAttendee)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeePushGuardianRS) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval attendee.BmAttendee = b.BrickInstance().Pr.(attendee.BmAttendee)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

