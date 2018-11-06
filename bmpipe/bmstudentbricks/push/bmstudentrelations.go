package studentpush

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

type BMStudentRSPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMStudentRSPushBrick) Exec() error {
	var tmp = b.bk.Pr.(student.BMStudent)
	eq := request.EQCond{}
	eq.Ky = "student_id"
	eq.Vy = tmp.Id

	var guardianIds []interface{}
	var contactIds []interface{}
	for _, v := range tmp.Guardians {
		v.InsertBMObject()
		guardianIds = append(guardianIds, v.Id)
	}
	//for _, v := range tmp.Contacts {
	//	v.InsertBMObject()
	//	contactIds = append(contactIds, v.Id)
	//}

	//設計邏輯還未明確
	//var continuedCoursesIds []string
	//var completedCoursesIds []string
	//for _,v := range tmp.ContinuedCourses {
	//	continuedCoursesIds = append(continuedCoursesIds, v.Id)
	//}
	//for _,v := range tmp.CompletedCourses {
	//	completedCoursesIds = append(completedCoursesIds, v.Id)
	//}

	req := request.Request{}
	req.Res = "BMStudentProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	var qr student.BMStudentProp
	err := qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.StudentID = tmp.Id
		qr.GuardianIds = guardianIds
		qr.ContactIds = contactIds
		//qr.ContinuedCoursesIds = continuedCoursesIds
		//qr.CompletedCoursesIds = completedCoursesIds
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	return nil
}

func (b *BMStudentRSPushBrick) Prepare(pr interface{}) error {
	req := pr.(student.BMStudent)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMStudentRSPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMStudentRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMStudentRSPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(student.BMStudent)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMStudentRSPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval student.BMStudent = b.BrickInstance().Pr.(student.BMStudent)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
