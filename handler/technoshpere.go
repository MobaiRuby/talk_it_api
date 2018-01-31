package handler

import (
	"net/http"
	"strconv"

	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/services/talk_it"
	"github.com/henrylee2cn/faygo"
)

func (h *Base) Technoshperes() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		ts, err := talk_it.Technoshperes(cur_user_id)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupDismissError,
				"msg":  err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "",
			"data": ts,
		})
	})
}
