package accountpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"io"
	"net/http"
)

type BMAccountPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAccountPushBrick) Exec() error {
	var err error
	var tmp account.BmAccount = b.bk.Pr.(account.BmAccount)

	if tmp.Id != "" && tmp.Id_.Valid() {
		if tmp.Valid() && tmp.IsAccountRegisted() {
			//TODO: error处理
			b.bk.Err = -8
		} else {
			//TODO: 临时版本暂时注掉rsa加密
			//TODO: 配置参数化
			//err = tmp.DecodeByCompanyDate("BlackMirror", "2018")
			//if err != nil {
			//	return err
			//}
			//tmp.Secret2MD5()

			tmp.InsertBMObject()
			tmp.SecretWord = ""
			b.bk.Pr = tmp
		}
	}

	return err
}

func (b *BMAccountPushBrick) Prepare(pr interface{}) error {

	req := pr.(account.BmAccount)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMAccountPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	ec := b.BrickInstance().Err
	if int(idx) < tmp-1 && ec == 0 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAccountPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAccountPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(account.BmAccount)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAccountPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		//reval := b.BrickInstance().Pr.(auth.BmLoginSucceedBySaaS)
		reval := b.BrickInstance().Pr.(account.BmAccount)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
