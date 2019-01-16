package sessionablepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmSessionablePushProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmSessionablePushProp) Exec() error {
	tmp := b.bk.Pr.(sessionable.BmSessionable)

	yard := tmp.Yard
	sbc := sessionable.BmSessionableBindYard{}
	sbc.SessionableId = tmp.Id
	sbc.YardId = yard.Id
	sbc.Id_ = bson.NewObjectId()
	sbc.Id = sbc.Id_.Hex()
	sbc.CheckExist()
	err := sbc.InsertBMObject()

	sessionInfo := tmp.SessionInfo
	sbs := sessionable.BmSessionableBindSessionInfo{}
	sbs.SessionableId = tmp.Id
	sbs.SessionInfoId = sessionInfo.Id
	sbs.Id_ = bson.NewObjectId()
	sbs.Id = sbs.Id_.Hex()
	sbs.CheckExist()
	err = sbs.InsertBMObject()

	sbt := sessionable.BmSessionableBindTeacher{}
	sbt.Id_ = bson.NewObjectId()
	sbt.Id = sbt.Id_.Hex()
	sbt.SessionableId = tmp.Id
	sbt.Clear()
	for _, item := range tmp.Teachers {

		//push classTeacher
		item.InsertBMObject()

		ist := sessionable.BmSessionableBindTeacher{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.SessionableId = tmp.Id
		ist.TeacherId = item.Id
		ist.InsertBMObject()
	}

	rby := sessionable.BmSessionableBindAttendee{}
	rby.Id_ = bson.NewObjectId()
	rby.Id = rby.Id_.Hex()
	rby.SessionableId = tmp.Id
	rby.Clear()
	for _, item := range tmp.Attendees {
		ist := sessionable.BmSessionableBindAttendee{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.SessionableId = tmp.Id
		ist.AttendeeId = item.Id
		ist.InsertBMObject()
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmSessionablePushProp) Prepare(pr interface{}) error {
	req := pr.(sessionable.BmSessionable)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmSessionablePushProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmSessionablePushProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmSessionablePushProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessionable.BmSessionable)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmSessionablePushProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval sessionable.BmSessionable = b.BrickInstance().Pr.(sessionable.BmSessionable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

