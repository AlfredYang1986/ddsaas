package main

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/ddsaas/bmmodel/account"
	"github.com/alfredyang1986/ddsaas/bmmodel/address"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"github.com/alfredyang1986/ddsaas/bmmodel/brand"
	"github.com/alfredyang1986/ddsaas/bmmodel/guardian"
	"github.com/alfredyang1986/ddsaas/bmmodel/payment"
	"github.com/alfredyang1986/ddsaas/bmmodel/person"
	"github.com/alfredyang1986/ddsaas/bmmodel/region"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmaccountbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmattendeebricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmattendeebricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmbrandbricks/push"
	"net/http"
	"sync"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/others"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/update"
	"github.com/alfredyang1986/ddsaas/bmpipe/bmteacherbricks/push"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("Request", &request.Request{})
	fac.RegisterModel("EqCond", &request.EqCond{})
	fac.RegisterModel("UpCond", &request.UpCond{})
	fac.RegisterModel("FmCond", &request.FmCond{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})

	fac.RegisterModel("BMRsaKey", &auth.BMRsaKey{})
	fac.RegisterModel("BMAccount", &account.BMAccount{})
	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWeChat", &auth.BMWeChat{})
	fac.RegisterModel("BMAuthProp", &auth.BMAuthProp{})
	fac.RegisterModel("BmAttendee", &attendee.BmAttendee{})
	fac.RegisterModel("BmAttendees", &attendee.BmAttendees{})
	fac.RegisterModel("BMAttendeeProp", &attendee.BMAttendeeProp{})
	fac.RegisterModel("BMAttendeeGuardianRS", &attendee.BMAttendeeGuardianRS{})
	fac.RegisterModel("BMAttendeeGuardianRSeS", &attendee.BMAttendeeGuardianRSeS{})
	fac.RegisterModel("BmGuardian", &guardian.BmGuardian{})
	fac.RegisterModel("BMGuardianProp", &guardian.BMGuardianProp{})
	fac.RegisterModel("BmPerson", &person.BmPerson{})
	fac.RegisterModel("BmAddress", &address.BmAddress{})
	fac.RegisterModel("BmRegion", &region.BmRegion{})
	fac.RegisterModel("BMPayment", &payment.BMPayment{})

	fac.RegisterModel("BMBrand", &brand.BMBrand{})
	fac.RegisterModel("BMTeacher", &teacher.BMTeacher{})

	/*------------------------------------------------
	 * auth find bricks object
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

	/*------------------------------------------------
	 * attendee push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAttendeePushBrick", &attendeepush.BMAttendeePushBrick{})
	fac.RegisterModel("BMAttendeePushPerson", &attendeepush.BMAttendeePushPerson{})
	fac.RegisterModel("BMAttendeePushPersonRS", &attendeepush.BMAttendeePushPersonRS{})
	fac.RegisterModel("BMAttendeePushGuardian", &attendeepush.BMAttendeePushGuardian{})
	fac.RegisterModel("BMAttendeePushGuardianRS", &attendeepush.BMAttendeePushGuardianRS{})

	fac.RegisterModel("BMBrandPushBrick", &brandpush.BMBrandPushBrick{})
	fac.RegisterModel("BMTeacherPushBrick", &teacherpush.BMTeacherPushBrick{})

	/*------------------------------------------------
	 * auth update bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneUpdateBrick", &authupdate.BMAuthPhoneUpdateBrick{})
	fac.RegisterModel("BMAuthWechatUpdateBrick", &authupdate.BMAuthWechatUpdateBrick{})

	/*------------------------------------------------
	 * auth delete bricks object
	 *------------------------------------------------*/

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthGenerateToken", &authothers.BMAuthGenerateToken{})

	r := bmrouter.BindRouter()

	var once sync.Once
	var bmRouter bmconfig.BMRouterConfig
	once.Do(bmRouter.GenerateConfig)

	http.ListenAndServe(":" + bmRouter.Port, r)
}
