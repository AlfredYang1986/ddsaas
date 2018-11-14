package reservablepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/reservable"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmBindReservableProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmBindReservableProp) Exec() error {
	tmp := b.bk.Pr.(reservable.BmReservable)

	session := tmp.SessionInfo
	sbc := reservable.BmReservableBindSession{}
	sbc.ReservableId = tmp.Id
	sbc.SessionId = session.Id
	sbc.Id_ = bson.NewObjectId()
	sbc.Id = sbc.Id_.Hex()
	sbc.CheckExist()
	err := sbc.InsertBMObject()

	rby := reservable.BmReservableBindYard{}
	rby.Id_ = bson.NewObjectId()
	rby.Id = rby.Id_.Hex()
	rby.ReservableId = tmp.Id
	rby.Clear()

	for _, item := range tmp.Yards {
		ist := reservable.BmReservableBindYard{}
		ist.Id_ = bson.NewObjectId()
		ist.Id = ist.Id_.Hex()
		ist.ReservableId = tmp.Id
		ist.YardId = item.Id
		ist.InsertBMObject()
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmBindReservableProp) Prepare(pr interface{}) error {
	req := pr.(reservable.BmReservable)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmBindReservableProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmBindReservableProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmBindReservableProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(reservable.BmReservable)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmBindReservableProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(reservable.BmReservable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

