package main

import (
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"github.com/alfredyang1986/ddsaas/bmmodel/activity"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmactivitybricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmbrandbricks/push"
	"net/http"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"github.com/alfredyang1986/ddsaas/bmmodel/class"
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
	"github.com/alfredyang1986/ddsaas/bmpipe/bmclassbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcontactbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcontactbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcoursebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmlocationbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/delete"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmorderbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmstudentbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmstudentbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/push"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("request", &request.Request{})
	fac.RegisterModel("eqcond", &request.EQCond{})
	fac.RegisterModel("upcond", &request.UPCond{})
	fac.RegisterModel("fmcond", &request.FMUCond{})

	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMRsaKey", &auth.BMRsaKey{})
	fac.RegisterModel("BMAccount", &account.BMAccount{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWeChat", &auth.BMWeChat{})
	fac.RegisterModel("BMAuthProp", &auth.BMAuthProp{})
	fac.RegisterModel("BMProfile", &profile.BMProfile{})
	fac.RegisterModel("BMProfileCompanyRS", &profile.BMProfileCompanyRS{})
	fac.RegisterModel("BMCompany", &profile.BMCompany{})
	fac.RegisterModel("BMBrand", &brand.BMBrand{})
	fac.RegisterModel("BMBrandLocationRS", &brand.BMBrandLocationRS{})
	fac.RegisterModel("BMLocation", &location.BMLocation{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})
	fac.RegisterModel("BMCourse", &course.BMCourse{})
	fac.RegisterModel("BMStudent", &student.BMStudent{})
	fac.RegisterModel("BMStudents", &student.BMStudents{})
	fac.RegisterModel("BMGuardian", &student.BMGuardian{})
	fac.RegisterModel("BMContacter", &student.BMContacter{})
	fac.RegisterModel("BMStudentProp", &student.BMStudentProp{})
	fac.RegisterModel("BMStudentsProp", &student.BMStudentsProp{})
	fac.RegisterModel("BMTeacher", &teacher.BMTeacher{})
	fac.RegisterModel("BMClass", &class.BMClass{})
	fac.RegisterModel("BMActivity", &activity.BMActivity{})
	fac.RegisterModel("BMActivityBrandRS", &activity.BMActivityBrandRS{})

	fac.RegisterModel("Contact", &contact.Contact{})             //del
	fac.RegisterModel("BMContactProp", &contact.BMContactProp{}) //del
	fac.RegisterModel("Order", &order.Order{})                   //del

	/*------------------------------------------------
	 * auth find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneFindBrick", &authfind.BMAuthPhoneFindBrick{})
	fac.RegisterModel("BMAuthRS2AuthBrick", &authfind.BMAuthRS2AuthBrick{})
	fac.RegisterModel("BMPhone2AuthRSBrick", &authfind.BMPhone2AuthRSBrick{})
	fac.RegisterModel("BMGetPublicKeyBrick", &authfind.BMGetPublicKeyBrick{})
	fac.RegisterModel("BMAccountFindBrick", &accountfind.BMAccountFindBrick{})

	fac.RegisterModel("BMStudentFindBrick", &studentfind.BMStudentFindBrick{})
	fac.RegisterModel("BMStudent2StudentRSBrick", &studentfind.BMStudent2StudentRSBrick{})
	fac.RegisterModel("BMStudentRS2StudentBrick", &studentfind.BMStudentRS2StudentBrick{})
	fac.RegisterModel("BMStudentFindMultiBrick", &studentfind.BMStudentFindMultiBrick{})
	fac.RegisterModel("BMStudents2StudentRSBrick", &studentfind.BMStudents2StudentRSBrick{})
	fac.RegisterModel("BMStudentRS2StudentsBrick", &studentfind.BMStudentRS2StudentsBrick{})

	fac.RegisterModel("BMContactFindBrick", &contactfind.BMContactFindBrick{})     //del
	fac.RegisterModel("BMOrderFindBrick", &orderfind.BMOrderFindBrick{})           //del
	fac.RegisterModel("BMOrderFindMultiBrick", &orderfind.BMOrderFindMultiBrick{}) //del

	/*------------------------------------------------
	 * auth push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMPhonePushBrick", &authpush.BMPhonePushBrick{})
	fac.RegisterModel("BMWechatPushBrick", &authpush.BMWechatPushBrick{})
	fac.RegisterModel("BMProfilePushBrick", &authpush.BMProfilePushBrick{})
	fac.RegisterModel("BMAuthCompanyPushBrick", &authpush.BMAuthCompanyPushBrick{})
	fac.RegisterModel("BMProfileCompanyRSPushBrick", &authpush.BMProfileCompanyRSPushBrick{})
	fac.RegisterModel("BMAuthRSPushBrick", &authpush.BMAuthRSPushBrick{})
	fac.RegisterModel("BMAuthPushBrick", &authpush.BMAuthPushBrick{})
	fac.RegisterModel("BMRsaKeyGenerateBrick", &authpush.BMRsaKeyGenerateBrick{})
	fac.RegisterModel("BMAccountPushBrick", &accountpush.BMAccountPushBrick{})

	fac.RegisterModel("BMBrandPushBrick", &brandpush.BMBrandPushBrick{})
	fac.RegisterModel("BMBrandPushLocationBrick", &brandpush.BMBrandPushLocationBrick{})
	fac.RegisterModel("BMBrandLocationRSPush", &brandpush.BMBrandLocationRSPush{})
	fac.RegisterModel("BMBrandCompanyRSPush", &brandpush.BMBrandCompanyRSPush{})
	fac.RegisterModel("BMLocationPushBrick", &locationpush.BMLocationPushBrick{})
	fac.RegisterModel("BMCoursePushBrick", &coursepush.BMCoursePushBrick{})
	fac.RegisterModel("BMStudentPushBrick", &studentpush.BMStudentPushBrick{})
	fac.RegisterModel("BMStudentRSPushBrick", &studentpush.BMStudentRSPushBrick{})
	fac.RegisterModel("BMTeacherPushBrick", &teacherpush.BMTeacherPushBrick{})
	fac.RegisterModel("BMClassPushBrick", &classpush.BMClassPushBrick{})
	fac.RegisterModel("BMActivityPushBrick", &activitypush.BMActivityPushBrick{})
	fac.RegisterModel("BMActivityBrandRSPush", &activitypush.BMActivityBrandRSPush{})

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
