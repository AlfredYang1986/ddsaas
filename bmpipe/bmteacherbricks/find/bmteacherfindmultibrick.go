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

type BmTeacherFindMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherFindMultiBrick) Exec() error {
	if b.bk.Req.Res == "BmTeacher" {
		var tmp teacher.BmTeachers
		err := tmp.FindMulti(*b.bk.Req)
		b.bk.Pr = tmp
		return err
	} else if b.bk.Req.Res == "BmPerson" {
		return errors.New("query Person")
	} else {
		return errors.New("query condition error !")
	}
}

func (b *BmTeacherFindMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmTeacherFindMultiBrick) Done(pkg string, idx int64, e error) error {

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

func (b *BmTeacherFindMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherFindMultiBrick) ResultTo(w io.Writer) error {

	if b.bk.Req.Res == "BmTeacher" {
		pr := b.BrickInstance().Pr
		tmp := pr.(teacher.BmTeachers)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		pr := b.BrickInstance().Pr
		tmp := pr.(request.Request)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *BmTeacherFindMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeachers = b.BrickInstance().Pr.(teacher.BmTeachers)
		jsonapi.ToJsonAPI(reval.Teachers, w)
	}
}