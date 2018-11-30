package main

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"github.com/alfredyang1986/ddsaas/bmmodel/address"
	"github.com/alfredyang1986/ddsaas/bmmodel/apply"
	"github.com/alfredyang1986/ddsaas/bmmodel/applyee"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"github.com/alfredyang1986/ddsaas/bmmodel/category"
	"github.com/alfredyang1986/ddsaas/bmmodel/certification"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/honor"
	"github.com/alfredyang1986/ddsaas/bmmodel/kid"
	"github.com/alfredyang1986/ddsaas/bmmodel/payment"
	"github.com/alfredyang1986/ddsaas/bmmodel/region"
	"github.com/alfredyang1986/ddsaas/bmmodel/reservable"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"github.com/alfredyang1986/ddsaas/bmmodel/wechatinfo"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmapplybricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmapplybricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmapplyeebricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmapplyeebricks/other"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmapplyeebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmattendeebricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmattendeebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmattendeebricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmbrandbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmbrandbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmbrandbricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcategorybricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcourseinfobricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmquerrywechatopenid/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmreservablebricks/delect"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmreservablebricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmreservablebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmsessionablebricks/delete"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmsessionablebricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmsessionablebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmtagimgsbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmyardbricks/update"
	"net/http"
	"sync"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"github.com/alfredyang1986/ddsaas/bmmodel/room"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessioninfo"
	"github.com/alfredyang1986/ddsaas/bmmodel/tagimg"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/ddsaas/bmmodel/yard"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/others"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcourseinfobricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcourseinfobricks/findmulti"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmcourseinfobricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmyardbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmyardbricks/push"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("Request", &request.Request{})
	fac.RegisterModel("Eqcond", &request.Eqcond{})
	fac.RegisterModel("Upcond", &request.Upcond{})
	fac.RegisterModel("Fmcond", &request.Fmcond{})
	fac.RegisterModel("BmErrorNode", &bmerror.BmErrorNode{})

	fac.RegisterModel("BMRsaKey", &auth.BMRsaKey{})
	fac.RegisterModel("BmAccount", &account.BmAccount{})
	fac.RegisterModel("BmBindAccountBrand", &account.BmBindAccountBrand{})
	fac.RegisterModel("BmAuth", &auth.BmAuth{})
	fac.RegisterModel("BmPhone", &auth.BmPhone{})
	fac.RegisterModel("BmWeChat", &auth.BmWeChat{})
	fac.RegisterModel("BmAuthProp", &auth.BmAuthProp{})
	fac.RegisterModel("BmLoginSucceed", &auth.BmLoginSucceed{})

	fac.RegisterModel("BmAttendee", &attendee.BmAttendee{})
	fac.RegisterModel("BmAttendees", &attendee.BmAttendees{})
	fac.RegisterModel("BMAttendeeProp", &attendee.BMAttendeeProp{})
	fac.RegisterModel("BMAttendeeGuardianRS", &attendee.BMAttendeeGuardianRS{})
	fac.RegisterModel("BMAttendeeGuardianRSeS", &attendee.BMAttendeeGuardianRSeS{})
	fac.RegisterModel("BmGuardian", &guardian.BmGuardian{})
	//fac.RegisterModel("BmPerson", &person.BmPerson{})
	//fac.RegisterModel("BmPersons", &person.BmPersons{})
	fac.RegisterModel("BmAddress", &address.BmAddress{})
	fac.RegisterModel("BmRegion", &region.BmRegion{})
	fac.RegisterModel("BMPayment", &payment.BMPayment{})

	fac.RegisterModel("BmBrand", &brand.BmBrand{})
	fac.RegisterModel("BmHonor", &honor.BmHonor{})
	fac.RegisterModel("BmCategory", &category.BmCategory{})
	fac.RegisterModel("BmCertification", &certification.BmCertification{})

	fac.RegisterModel("BmTeacher", &teacher.BmTeacher{})
	fac.RegisterModel("BmTeachers", &teacher.BmTeachers{})

	/*------------------------------------------------
	 * find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneFindBrick", &authfind.BMAuthPhoneFindBrick{})
	fac.RegisterModel("BMAuthRS2AuthBrick", &authfind.BMAuthRS2AuthBrick{})
	fac.RegisterModel("BMPhone2AuthRSBrick", &authfind.BMPhone2AuthRSBrick{})
	fac.RegisterModel("BMGetPublicKeyBrick", &authfind.BMGetPublicKeyBrick{})
	fac.RegisterModel("BMAccountFindBrick", &accountfind.BMAccountFindBrick{})

	fac.RegisterModel("BMAttendeeFindBrick", &attendeefind.BMAttendeeFindBrick{})
	fac.RegisterModel("BMAttendeeFindMulti", &attendeefind.BMAttendeeFindMulti{})
	fac.RegisterModel("BMAttendeeRS2Attendee", &attendeefind.BMAttendeeRS2Attendee{})

	/*------------------------------------------------
	 * auth push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMPhonePushBrick", &authpush.BMPhonePushBrick{})
	fac.RegisterModel("BMWechatPushBrick", &authpush.BMWechatPushBrick{})
	fac.RegisterModel("BMAuthRSPushBrick", &authpush.BMAuthRSPushBrick{})
	fac.RegisterModel("BMAuthPushBrick", &authpush.BMAuthPushBrick{})
	fac.RegisterModel("BMRsaKeyGenerateBrick", &authpush.BMRsaKeyGenerateBrick{})
	fac.RegisterModel("BMAccountPushBrick", &accountpush.BMAccountPushBrick{})
	fac.RegisterModel("BmAccountBindBrand", &accountpush.BmAccountBindBrand{})

	/*------------------------------------------------
	 * attendee push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAttendeePushBrick", &attendeepush.BMAttendeePushBrick{})
	fac.RegisterModel("BMAttendeePushGuardian", &attendeepush.BMAttendeePushGuardian{})
	fac.RegisterModel("BMAttendeePushGuardianRS", &attendeepush.BMAttendeePushGuardianRS{})

	/*------------------------------------------------
	 * brand bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BmBrandPushBrick", &brandpush.BmBrandPushBrick{})
	fac.RegisterModel("BmBrandPushProp", &brandpush.BmBrandPushProp{})
	fac.RegisterModel("BmBrandBindProp", &brandpush.BmBrandBindProp{})
	fac.RegisterModel("BmBrandFindBrick", &brandfind.BmBrandFindBrick{})

	fac.RegisterModel("BmTeacherPushBrick", &teacherpush.BmTeacherPushBrick{})
	//fac.RegisterModel("BmTeacherPersonPushBrick", &teacherpush.BmTeacherPersonPushBrick{})
	//fac.RegisterModel("BmTeacherPushPersonRS", &teacherpush.BmTeacherPushPersonRS{})
	fac.RegisterModel("BmTeacherFindBrick", &teacherfind.BmTeacherFindBrick{})
	//fac.RegisterModel("BmTeacherRS2Teacher", &teacherfind.BmTeacherRS2Teacher{})
	//fac.RegisterModel("BmPersonFindBrick", &teather_person.BmPersonTeacherFindBrick{})
	//fac.RegisterModel("BmPersonTeacherRS", &teather_person.BmPersonTeacherRS{})
	fac.RegisterModel("BmTeacherFindMultiBrick", &teacherfind.BmTeacherFindMultiBrick{})
	//fac.RegisterModel("BmTeacherMultiRS", &teacherfind.BmTeacherMultiRS{})

	/*------------------------------------------------
	 * yard brick object
	 *------------------------------------------------*/
	fac.RegisterModel("BmYard", &yard.BmYard{})
	fac.RegisterModel("BmBindYardImg", &yard.BmBindYardImg{})
	fac.RegisterModel("BmYards", &yard.BmYards{})
	fac.RegisterModel("BmRoom", &room.BmRoom{})
	fac.RegisterModel("BmTagImg", &tagimg.BmTagImg{})

	fac.RegisterModel("BmYardPushBrick", &yardpush.BmYardPushBrick{})
	fac.RegisterModel("BmTagImgYardPushBrick", &yardpush.BmTagImgYardPushBrick{})
	fac.RegisterModel("BmYardRoomPushBrick", &yardpush.BmYardRoomPushBrick{})
	fac.RegisterModel("BmYardPushCertificationBrick", &yardpush.BmYardPushCertificationBrick{})
	fac.RegisterModel("BmBindYardPropBrick", &yardpush.BmBindYardPropBrick{})
	fac.RegisterModel("BmYardFindBrick", &yardfind.BmYardFindBrick{})
	fac.RegisterModel("BmYardFindMulti", &yardfind.BmYardFindMulti{})
	fac.RegisterModel("BmTagImgPushBrick", &tagimgpush.BmTagImgPushBrick{})
	fac.RegisterModel("BmTagImgBindYard", &tagimgpush.BmTagImgBindYard{})

	/*------------------------------------------------
	 * session brick object
	 *------------------------------------------------*/
	fac.RegisterModel("BmSessionInfo", &sessioninfo.BmSessionInfo{})
	fac.RegisterModel("BmSessionInfos", &sessioninfo.BmSessionInfos{})

	fac.RegisterModel("BmSessionInfoPushBrick", &courseinfopush.BmSessionInfoPushBrick{})
	fac.RegisterModel("BmSessionCatPushBrick", &courseinfopush.BmSessionCatPushBrick{})
	fac.RegisterModel("BmSessionImgPushBrick", &courseinfopush.BmSessionImgPushBrick{})
	fac.RegisterModel("BmSessionPushProp", &courseinfopush.BmSessionPushProp{})
	fac.RegisterModel("BmFindSessionInfoBrick", &courseinfofind.BmFindSessionInfoBrick{})
	fac.RegisterModel("BmFindSessionInfoMultiBrick", &courseinfofindmulti.BmFindSessionInfoMultiBrick{})

	/*------------------------------------------------
	 * reservable brick object
	 *------------------------------------------------*/
	fac.RegisterModel("BmReservable", &reservable.BmReservable{})
	fac.RegisterModel("BmReservableBindSession", &reservable.BmReservableBindSession{})

	fac.RegisterModel("BmReservablePushBrick", &reservablepush.BmReservablePushBrick{})
	fac.RegisterModel("BmReservablePushSession", &reservablepush.BmReservablePushSession{})
	fac.RegisterModel("BmBindReservableProp", &reservablepush.BmBindReservableProp{})
	fac.RegisterModel("BmReservableFindBrick", &reservablefind.BmReservableFindBrick{})
	fac.RegisterModel("BmReservableDeleteBrick", &reservabledelete.BmReservableDeleteBrick{})
	fac.RegisterModel("BmReservableFindMulti", &reservablefind.BmReservableFindMulti{})

	/*------------------------------------------------
	 * sessionable brick object
	 *------------------------------------------------*/
	fac.RegisterModel("BmSessionable", &sessionable.BmSessionable{})
	fac.RegisterModel("BmSessionableBindYard", &sessionable.BmSessionableBindYard{})
	fac.RegisterModel("BmSessionableBindTeacher", &sessionable.BmSessionableBindTeacher{})
	fac.RegisterModel("BmSessionableBindAttendee", &sessionable.BmSessionableBindAttendee{})

	fac.RegisterModel("BmSessionablePushBrick", &sessionablepush.BmSessionablePushBrick{})
	fac.RegisterModel("BmSessionablePushProp", &sessionablepush.BmSessionablePushProp{})
	fac.RegisterModel("BmSessionableFindBrick", &sessionablefind.BmSessionableFindBrick{})
	fac.RegisterModel("BmSessionableDeleteBrick", &sessionabledelete.BmSessionableDeleteBrick{})
	fac.RegisterModel("BmSessionableFindMulti", &sessionablefind.BmSessionableFindMulti{})

	/*------------------------------------------------
	 * apply brick object
	 *------------------------------------------------*/
	fac.RegisterModel("BmKid", &kid.BmKid{})
	fac.RegisterModel("BmApply", &apply.BmApply{})
	fac.RegisterModel("BmApplyee", &applyee.BmApplyee{})
	fac.RegisterModel("BmApplyBindKid", &apply.BmApplyBindKid{})
	fac.RegisterModel("BmApplyBindApplyee", &apply.BmApplyBindApplyee{})

	fac.RegisterModel("BmApplyPushBrick", &applypush.BmApplyPushBrick{})
	fac.RegisterModel("BmApplyPushKids", &applypush.BmApplyPushKids{})
	fac.RegisterModel("BmApplyPushProp", &applypush.BmApplyPushProp{})

	fac.RegisterModel("BmApplyFindBrick", &applyfind.BmApplyFindBrick{})
	fac.RegisterModel("BmAppliesFindBrick", &applyfind.BmAppliesFindBrick{})

	fac.RegisterModel("BmApplyeePushBrick", &applyeepush.BmApplyeePushBrick{})
	fac.RegisterModel("BmApplyeeFindBrick", &applyeefind.BmApplyeeFindBrick{})
	fac.RegisterModel("BmApplyeeGenerateToken", &applyeeother.BmApplyeeGenerateToken{})

	/*------------------------------------------------
	 * update bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneUpdateBrick", &authupdate.BMAuthPhoneUpdateBrick{})
	fac.RegisterModel("BMAuthWechatUpdateBrick", &authupdate.BMAuthWechatUpdateBrick{})
	fac.RegisterModel("BmAttendeeUpdateBrick", &attendeeupdate.BmAttendeeUpdateBrick{})
	fac.RegisterModel("BmGuardianUpdateBrick", &attendeeupdate.BmGuardianUpdateBrick{})
	fac.RegisterModel("BmTeacherUpdateBrick", &teacherupdate.BmTeacherUpdateBrick{})
	fac.RegisterModel("BmYardUpdateBrick", &yardupdate.BmYardUpdateBrick{})
	fac.RegisterModel("BmSessionInfoUpdateBrick", &courseinfoupdate.BmSessionInfoUpdateBrick{})
	fac.RegisterModel("BmBrandUpdateBrick", &brandupdate.BmBrandUpdateBrick{})
	fac.RegisterModel("BmCategoryUpdateBrick", &categoryupdate.BmCategoryUpdateBrick{})

	fac.RegisterModel("BmWeChatInfo", &wechatinfo.BmWeChatInfo{})
	fac.RegisterModel("BmWeChatInfoFindBrick", &wechatopenidfind.BmWeChatInfoFindBrick{})

	/*------------------------------------------------
	 * delete bricks object
	 *------------------------------------------------*/

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthGenerateToken", &authothers.BMAuthGenerateToken{})

	r := bmrouter.BindRouter()

	var once sync.Once
	var bmRouter bmconfig.BMRouterConfig
	once.Do(bmRouter.GenerateConfig)

	http.ListenAndServe(":"+bmRouter.Port, r)
}
