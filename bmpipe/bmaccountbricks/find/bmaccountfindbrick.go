package accountfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
)

type BMAccountFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAccountFindBrick) Exec() error {
	var tmp account.BmAccount
	err := tmp.FindOne(*b.bk.Req)
	tmp.SecretWord = ""
	b.bk.Pr = tmp
	return err
}

func (b *BMAccountFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)

	//var eqCondArr []request.Eqcond
	//for _, e := range req.Eqcond {
	//	if e.Ky == "secretword" {
	//		tmpAccount := account.BmAccount{
	//			SecretWord: e.Vy.(string),
	//		}
	//		//TODO: 配置参数化
	//		tmpAccount.DecodeByCompanyDate("BlackMirror", "2018")
	//		tmpAccount.Secret2MD5()
	//		e.Vy = tmpAccount.SecretWord
	//	}
	//	eqCondArr = append(eqCondArr, e)
	//}
	//req.Eqcond = eqCondArr

	b.BrickInstance().Req = &req
	return nil
}

func (b *BMAccountFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAccountFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAccountFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(account.BmAccount)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAccountFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval account.BmAccount = b.BrickInstance().Pr.(account.BmAccount)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
