package courseunitfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/courseunit"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
)

type BmCourseUnitFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmCourseUnitFindBrick) Exec() error {
	var tmp courseunit.BmCourseUnit
	err := tmp.FindOne(*b.bk.Req)
	tmp.ReSetProp()
	b.bk.Pr = tmp
	return err
}

func (b *BmCourseUnitFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmCourseUnitFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmCourseUnitFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmCourseUnitFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(courseunit.BmCourseUnit)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmCourseUnitFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval courseunit.BmCourseUnit = b.BrickInstance().Pr.(courseunit.BmCourseUnit)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

