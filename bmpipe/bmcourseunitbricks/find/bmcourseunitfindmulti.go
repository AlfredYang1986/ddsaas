package courseunitfind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/courseunit"
	"io"
	"net/http"
)

type BmCourseUnitFindMulti struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmCourseUnitFindMulti) Exec() error {
	var tmp courseunit.BmCourseUnits
	err := tmp.FindMulti(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BmCourseUnitFindMulti) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmCourseUnitFindMulti) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmCourseUnitFindMulti) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmCourseUnitFindMulti) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(courseunit.BmCourseUnits)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmCourseUnitFindMulti) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		results := b.BrickInstance().Pr.(courseunit.BmCourseUnits)
		jsonapi.ToJsonAPI(results.CourseUnits, w)
	}
}
