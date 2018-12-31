package roompush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BmRoomPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmRoomPushBrick) Exec() error {
	tmp := b.bk.Pr.(room.BmRoom)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *BmRoomPushBrick) Prepare(pr interface{}) error {
	req := pr.(room.BmRoom)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmRoomPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmRoomPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmRoomPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(room.BmRoom)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmRoomPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval room.BmRoom = b.BrickInstance().Pr.(room.BmRoom)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
