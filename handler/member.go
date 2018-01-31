package handler

import (
	"net/http"
	"strconv"

	"fmt"

	"github.com/MobaiRuby/talk_it_api/cache"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/models"
	"github.com/MobaiRuby/talk_it_api/services/talk_it"
	"github.com/henrylee2cn/faygo"
)

func (h *Base) Members() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		ms, err := talk_it.Members(cur_user_id)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeMembersError,
				"msg":  err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "",
			"data": ms,
		})
	})
}

func (h *Base) AddMember() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		member := &models.Member{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), member); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeMembersAddError,
				"msg":  err.Error(),
			})
		}

		member.MasterId = cur_user_id
		if err := member.Add(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeMembersAddError,
				"msg":  err.Error(),
			})
		}

		cache.GetMemberCache().Store(member.Id, member)
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  fmt.Sprintf("add member to group %s success", member.GroupName),
		})
	})
}

func (h *Base) RemoveMember() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		member := &models.Member{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), member); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeMembersRemoveError,
				"msg":  err.Error(),
			})
		}

		member.MasterId = cur_user_id
		if err := member.Remove(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeMembersRemoveError,
				"msg":  err.Error(),
			})
		}

		cache.GetMemberCache().Delete(member.Id)
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  fmt.Sprintf("remove member from group %s success", member.GroupName),
		})
	})
}

func (h *Base) LeaveGroup() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		member := &models.Member{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), member); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupLeaveError,
				"msg":  err.Error(),
			})
		}

		member.UserId = cur_user_id
		if err := member.Leave(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeGroupLeaveError,
				"msg":  err.Error(),
			})
		}

		cache.GetMemberCache().Delete(member.Id)
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "leave group success",
		})
	})
}
