package studentfind

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/student"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMStudentRS2StudentBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudentRS2StudentBrick) Exec() error {
	prop := b.bk.Pr.(student.BMStudentProp)
	reval, err := findStudent(prop)
	guard, err := findGuardians(prop)
	conta, err := findContacts(prop)
	reval.Guardians = guard
	reval.Contacts = conta
	b.bk.Pr = reval
	return err
}

func (b *BMStudentRS2StudentBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudentProp)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudentRS2StudentBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudentRS2StudentBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudentRS2StudentBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudent)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudentRS2StudentBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudent = b.BrickInstance().Pr.(student.BMStudent)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func findStudent(prop student.BMStudentProp) (student.BMStudent, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.StudentID)
	req := request.Request{}
	req.Res = "BMStudent"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := student.BMStudent{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}

func findGuardians(prop student.BMStudentProp) ([]student.BMGuardian, error) {

	var gs []student.BMGuardian

	for _, v := range prop.GuardianIds {
		eq := request.EQCond{}
		eq.Ky = "_id"
		eq.Vy = bson.ObjectIdHex(v.(string))
		req := request.Request{}
		req.Res = "BMGuardian"
		var condi []interface{}
		condi = append(condi, eq)
		c := req.SetConnect("conditions", condi)
		fmt.Println(c)

		reval := student.BMGuardian{}
		err := reval.FindOne(c.(request.Request))
		if err != nil {
			return nil, err
		}
		gs = append(gs, reval)
	}

	return gs, nil
}

func findContacts(prop student.BMStudentProp) ([]student.BMContacter, error) {

	var cs []student.BMContacter

	for _, v := range prop.ContactIds {
		eq := request.EQCond{}
		eq.Ky = "_id"
		eq.Vy = bson.ObjectIdHex(v.(string))
		req := request.Request{}
		req.Res = "BMContacter"
		var condi []interface{}
		condi = append(condi, eq)
		c := req.SetConnect("conditions", condi)
		fmt.Println(c)

		reval := student.BMContacter{}
		err := reval.FindOne(c.(request.Request))
		if err != nil {
			return nil, err
		}
		cs = append(cs, reval)
	}

	return cs, nil
}
