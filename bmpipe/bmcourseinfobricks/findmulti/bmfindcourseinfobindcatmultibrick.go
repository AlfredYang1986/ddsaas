package courseinfofindmulti

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmFindSessionInfoBindCatMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmFindSessionInfoBindCatMultiBrick) Exec() error {
	tmp := b.BrickInstance().Pr.(sessioninfo.BmSessionInfos)

	var sessionIds []string
	for _, item := range tmp.Sessions {
		sessionIds = append(sessionIds, item.Id)
	}
	multiCondBind := make(map[string]interface{})
	multiCondBind["$in"] = sessionIds

	var sessionCatBinds []sessioninfo.BmSessionBindCat
	or_condi := bson.M{"sessionId": multiCondBind}
	err := bmmodel.FindMutilWithBson("BmSessionBindCat", or_condi, &sessionCatBinds)

	var condi0 []interface{}
	for _, item := range sessionCatBinds {
		condi0 = append(condi0, bson.ObjectIdHex(item.CategoryId))
	}
	multiCatCond := make(map[string]interface{})
	multiCatCond["$in"] = condi0

	var cats []category.BmCategory
	or_condi0 := bson.M{"_id": multiCatCond}
	err = bmmodel.FindMutilWithBson("BmCategory", or_condi0, &cats)

	for i, c := range cats {
		c.ResetIdWithId_()
		cats[i] = c
	}

	for i, session := range tmp.Sessions {
		sid := session.Id
		var cat category.BmCategory
		for _, sbc := range sessionCatBinds {
			if sbc.SessionId == sid {
				for _, ct := range cats {
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