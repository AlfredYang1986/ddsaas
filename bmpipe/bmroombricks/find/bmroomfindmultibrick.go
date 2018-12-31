package roomfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"io"
	"net/http"
)

type BmRoomFindMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmRoomFindMultiBrick) Exec() error {
	var tmp room.BmRooms
	err := tmp.FindMulti(b.BrickInstance().Pr.(request.Request))
	b.bk.Pr = tmp
	return err
}

func (b *BmRoomFindMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmRoomFindMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmRoomFindMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmRoomFindMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(room.BmRooms)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmRoomFindMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(room.BmRooms)
		jsonapi.ToJsonAPI(reval.Rooms, w)
	}
}