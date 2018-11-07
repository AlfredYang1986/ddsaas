package courseinfofindmulti

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
)

type BmFindSessionInfoBindCatMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmFindSessionInfoBindCatMultiBrick) Exec() error {
	tmp := b.BrickInstance().Pr.(sessioninfo.BmSessionInfos)
	//err := tmp.FindMulti(b.BrickInstance().Pr.(request.Request))

	req := request.Request{}
	req.Res = "BmSessionBindCat"
	var condi []interface{}
	for _, item := range tmp.Sessions {
		eq := request.Eqcond{}
		eq.Ky = "sessionId"
		eq.Vy = item.Id
		condi = append(condi, eq)
	}
	c := req.SetConnect("conditions", condi)

	var reval sessioninfo.BmSessionInfoBindCats
	err := reval.FindMulti(c.(request.Request))

	req0 := request.Request{}
	req0.Res = "BmCategory"
	var condi0 []interface{}
	for _, item := range reval.Binds {
		eq := request.Eqcond{}
		eq.Ky = "id"
		eq.Vy = item.CategoryId
		condi0 = append(condi0, eq)
	}
	c0 := req0.SetConnect("conditions", condi0)

	var cats category.BmCategories
	err = cats.FindMulti(c0.(request.Request))

	for i, session := range tmp.Sessions {
		sid := session.Id
		var cat category.BmCategory
		for _, sbc := range reval.Binds {
			if sbc.SessionId == sid {
				for _, ct := range cats.Cats {
					if ct.Id == sbc.CategoryId {
						cat = ct
					}
				}
			}
		}
		session.Cat = cat
		tmp.Sessions[i] = session
	}

	b.bk.Pr = tmp
	return err
}

func (b *BmFindSessionInfoBindCatMultiBrick) Prepare(pr interface{}) error {
	req := pr.(sessioninfo.BmSessionInfos)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmFindSessionInfoBindCatMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmFindSessionInfoBindCatMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmFindSessionInfoBindCatMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessioninfo.BmSessionInfos)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmFindSessionInfoBindCatMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		reval := b.BrickInstance().Pr.(sessioninfo.BmSessionInfos)
		jsonapi.ToJsonAPI(&reval, w)
	}
}