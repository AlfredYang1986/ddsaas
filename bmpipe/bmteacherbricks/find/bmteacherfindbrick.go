package teacherfind

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"errors"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
)

type BmTeacherFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherFindBrick) Exec() error {
	if b.bk.Req.Res == "BmTeacher" {
		var tmp teacher.BmTeacher
		err := tmp.FindOne(*b.bk.Req)
		b.bk.Pr = tmp
		return err
	} else if b.bk.Req.Res == "BmPerson" {
		return errors.New("query Person")
	} else {
		return errors.New("query condition error !")
	}
}

func (b *BmTeacherFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmTeacherFindBrick) Done(pkg string, idx int64, e error) error {

	if e != nil && e.Error() == "query Person" {
		b.BrickInstance().Pr = *b.bk.Req
		bmrouter.NextBrickRemote("findteacherprimary", 0, b)
	} else {
		tmp, _ := bmpkg.GetPkgLen(pkg)
		if int(idx) < tmp-1 {
			bmrouter.NextBrickRemote(pkg, idx+1, b)
		}
	}
	return nil
}

func (b *BmTeacherFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherFindBrick) ResultTo(w io.Writer) error {

	if b.bk.Req.Res == "BmTeacher" {
		pr := b.BrickInstance().Pr
		tmp := pr.(teacher.BmTeacher)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		pr := b.BrickInstance().Pr
		tmp := pr.(request.Request)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *BmTeacherFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		if b.bk.Req.Res == "BmTeacher" {
			pr := b.BrickInstance().Pr
			tmp := pr.(teacher.BmTeacher)
			jsonapi.ToJsonAPI(&tmp, w)
		} else if b.bk.Req.Res == "BmPerson" {
			var reval teacher.BmTeachers = b.BrickInstance().Pr.(teacher.BmTeachers)
			//jsonapi.ToJsonAPI(&reval, w)
			if len(reval.Teachers) > 0 {
				jsonapi.ToJsonAPI(&reval.Teachers[0], w)
			} else {
				bmerror.ErrInstance().ErrorReval(-9999, w)
			}
		}
	}
}
