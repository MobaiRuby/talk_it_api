package handler

import (
	"net/http"
	"strconv"

	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/models"
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
				"code": errorCodeTechnoshperesError,
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

func (h *Base) Technoshpere() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereError,
				"msg":  "get current user failed",
			})
		}
		t_id, err := pathId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereError,
				"msg":  err.Error(),
			})
		}

		if t, err := talk_it.Technoshpere(t_id); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereError,
				"msg":  err.Error(),
			})
		} else {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": 0,
				"msg":  "",
				"data": t,
			})
		}
	})
}

func (h *Base) AddTechnoshpere() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		ts := &models.Technoshpere{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), ts); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereAddError,
				"msg":  "get current user failed",
			})
		}

		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "add technoshpere success",
		})
	})
}

func (h *Base) RemoveTechnoshpere() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}
		t_id, err := pathId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereError,
				"msg":  err.Error(),
			})
		}

		t := models.Technoshpere{Id: t_id}
		if err := t.Remove(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeTechnoshpereRemoveError,
				"msg":  err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "remove technoshpere success",
		})
	})
}
