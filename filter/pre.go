package filter

import (
	"io/ioutil"

	"net/http"

	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/henrylee2cn/faygo"
)

var (
	noAuthPath []string
)

var Pre = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	// body放入ctx里, 用于log middleware
	bytes, _ := ioutil.ReadAll(ctx.R.Body)
	ctx.SetData(global.CTX_BODY, bytes)

	return nil
})

var Validate = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	if isNoAuth(ctx.Path()) {
		return nil
	}

	user_id := ctx.HeaderParam(global.TALK_IT_HEADER_CURRENT_USER_ID)
	if user_id == "" {
		return ctx.JSON(http.StatusUnauthorized, faygo.Map{
			"msg": "Not Allowed Request Without TalkIT User Id",
		})
	}
	ctx.SetData(global.TALK_IT_HEADER_CURRENT_USER_ID, user_id)

	return nil
})

func isNoAuth(path string) bool {
	if len(noAuthPath) == 0 {
		noAuthPath = []string{"/", "/login/tel", "login/wechat"}
	}
	for _, no_auth_path := range noAuthPath {
		if no_auth_path == path {
			return true
		}
	}
	return false
}
