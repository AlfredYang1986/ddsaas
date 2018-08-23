package main

import (
	"net/http"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"github.com/alfredyang1986/ddsaas/bmmodel/contact"
	"github.com/alfredyang1986/ddsaas/bmmodel/course"
	"github.com/alfredyang1986/ddsaas/bmmodel/location"
	"github.com/alfredyang1986/ddsaas/bmmodel/order"
	"github.com/alfredyang1986/ddsaas/bmmodel/profile"
	"github.com/alfredyang1986/ddsaas/bmmodel/student"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/others"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcontactbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcontactbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcoursebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmlocationbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/delete"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmstudentbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/push"
	"github.com/alfredyang1986/ddsaas/bmmodel/class"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmclassbricks/push"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWechat", &auth.BMWechat{})
	fac.RegisterModel("BMAuthProp", &auth.BMAuthProp{})
	fac.RegisterModel("BMProfile", &profile.BMProfile{})
	fac.RegisterModel("BMCompany", &profile.BMCompany{})
	fac.RegisterModel("BMLocation", &location.BMLocation{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})
	fac.RegisterModel("request", &request.Request{})
	fac.RegisterModel("eq_cond", &request.EQCond{})
	fac.RegisterModel("up_cond", &request.UPCond{})
	fac.RegisterModel("fm_cond", &request.FMUCond{})
	fac.RegisterModel("BMCourse", &course.BMCourse{})
	fac.RegisterModel("BMStudent", &student.BMStudent{})
	fac.RegisterModel("BMGuardian", &student.BMGuardian{})
	fac.RegisterModel("BMContacter", &student.BMContacter{})
	fac.RegisterModel("BMStudentProp", &student.BMStudentProp{})
	fac.RegisterModel("BMTeacher", &teacher.BMTeacher{})
	fac.RegisterModel("BMClass", &class.BMClass{})

	fac.RegisterModel("Contact", &contact.Contact{})             //del
	fac.RegisterModel("BMContactProp", &contact.BMContactProp{}) //del
	fac.RegisterModel("Order", &order.Order{})                   //del

	/*------------------------------------------------
	 * auth find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneFindBrick", &authfind.BMAuthPhoneFindBrick{})
	fac.RegisterModel("BMAuthRS2AuthBrick", &authfind.BMAuthRS2AuthBrick{})
	fac.RegisterModel("BMPhone2AuthRSBrick", &authfind.BMPhone2AuthRSBrick{})

	fac.RegisterModel("BMContactFindBrick", &contactfind.BMContactFindBrick{})     //del
	fac.RegisterModel("BMOrderFindBrick", &orderfind.BMOrderFindBrick{})           //del
	fac.RegisterModel("BMOrderFindMultiBrick", &orderfind.BMOrderFindMultiBrick{}) //del

	/*------------------------------------------------
	 * auth push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMPhonePushBrick", &authpush.BMPhonePushBrick{})
	fac.RegisterModel("BMWechatPushBrick", &authpush.BMWechatPushBrick{})
	fac.RegisterModel("BMProfilePushBrick", &authpush.BMProfilePushBrick{})
	fac.RegisterModel("BMAuthRSPushBrick", &authpush.BMAuthRSPushBrick{})
	fac.RegisterModel("BMAuthPushBrick", &authpush.BMAuthPushBrick{})

	fac.RegisterModel("BMLocationPushBrick", &locationpush.BMLocationPushBrick{})
	fac.RegisterModel("BMCoursePushBrick", &coursepush.BMCoursePushBrick{})
	fac.RegisterModel("BMStudentPushBrick", &studentpush.BMStudentPushBrick{})
	fac.RegisterModel("BMStudentRSPushBrick", &studentpush.BMStudentRSPushBrick{})
	fac.RegisterModel("BMTeacherPushBrick", &teacherpush.BMTeacherPushBrick{})
	fac.RegisterModel("BMClassPushBrick", &classpush.BMClassPushBrick{})

	fac.RegisterModel("BMOrderPushBrick", &orderpush.BMOrderPushBrick{})           //del
	fac.RegisterModel("BMContactPushBrick", &contactpush.BMContactPushBrick{})     //del
	fac.RegisterModel("BMContactRSPushBrick", &contactpush.BMContactRSPushBrick{}) //del

	/*------------------------------------------------
	 * auth update bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneUpdateBrick", &authupdate.BMAuthPhoneUpdateBrick{})
	fac.RegisterModel("BMAuthWechatUpdateBrick", &authupdate.BMAuthWechatUpdateBrick{})

	/*------------------------------------------------
	 * auth delete bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMOrderDeleteBrick", &orderdelete.BMOrderDeleteBrick{}) //del

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthGenerateToken", &authothers.BMAuthGenerateToken{})

	r := bmrouter.BindRouter()
	http.ListenAndServe(":8080", r)
}
