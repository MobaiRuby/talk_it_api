package handler

import (
	"net/http"

	"strconv"

	"github.com/MobaiRuby/talk_it_api/cache"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/models"
	"github.com/MobaiRuby/talk_it_api/services/talk_it"
	"github.com/henrylee2cn/faygo"
)

func (h *Base) Groups() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}
		groups, err := talk_it.Groups(cur_user_id)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupsError,
				"msg":  err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "",
			"data": groups,
		})
	})
}

func (h *Base) Group() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}
		g_id, err := pathId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFriendError,
				"msg":  err.Error(),
			})
		}

		if group, err := talk_it.Group(g_id); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFriendError,
				"msg":  err.Error(),
			})
		} else {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": 0,
				"msg":  "",
				"data": group,
			})
		}
	})
}

func (h *Base) CreateGroup() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		group := &models.Group{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), group); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupCreateError,
				"msg":  err.Error(),
			})
		}

		group.UserId = cur_user_id
		if err := group.Create(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupCreateError,
				"msg":  err.Error(),
			})
		}
		cache.GetGroupCache().Store(group.Id, group)
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "create group success",
		})
	})
}

func (h *Base) DismissGroup() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		group := &models.Group{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), group); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupDismissError,
				"msg":  err.Error(),
			})
		}

		group.UserId = cur_user_id
		if err := group.Dismiss(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupDismissError,
				"msg":  err.Error(),
			})
		}

		cache.GetGroupCache().Delete(group.Id)
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "dismiss group success",
		})
	})
}
