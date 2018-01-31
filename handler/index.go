package handler

import (
	"net/http"

	"fmt"

	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/services/facebook"
	"github.com/MobaiRuby/talk_it_api/services/talk_it"
	"github.com/henrylee2cn/faygo"
)

type (
	Admin struct{} // 后台api
	Base  struct{} // 通用api
	Fe    struct{} // 前台api
)

/*
** Index: 健康检查
 */
var Index = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	return ctx.JSON(http.StatusOK, faygo.Map{"info": "healthy"}, true)
})

// 手机号登录
func (h *Base) TelLogin() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		// step0: ctx bind
		auth := map[string]string{
			fbAuthCode: "",
		}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY), &auth); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeParamsError,
				"msg":  err.Error(),
			})
		}
		if _, ok := auth[fbAuthCode]; !ok {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeParamsError,
				"msg":  fmt.Sprintf("%s is blank", fbAuthCode),
			})
		}

		// step1: 接受fb_auth_code去fb获取fb_token
		fb_auth_token, err := facebook.GetToken(auth[fbAuthCode])
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFbTokenError,
				"msg":  err.Error(),
			})
		}
		if fb_auth_token == "" {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFbTokenError,
				"msg":  fmt.Errorf("get %s failed.", fbAuthToken),
			})
		}

		// step2: 用fb_token去fb获取fb_user
		v, err := facebook.ValidateToken(fb_auth_token)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFbValidationError,
				"msg":  err.Error(),
			})
		}
		if v == nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFbTokenError,
				"msg":  fmt.Errorf("fb token validation failed"),
			})
		}

		// step3: 用fb_user中的手机号去gb获取token(旧用户跳过step3)
		user, err := talk_it.Register2TalkIt(v.Phone.Number)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeRegisterUserError,
				"msg":  err.Error(),
			})
		}
		if user == nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeRegisterUserError,
				"msg":  fmt.Errorf("login failed"),
			})
		}

		// step4: 返回step3所获token和登录用户信息
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "login success",
			"data": user,
		})
	})
}

// 微信登录
func (h *Base) WeChatLogin() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "login success",
		})
	})
}
