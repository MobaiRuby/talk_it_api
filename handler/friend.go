package handler

import (
	"net/http"

	"strconv"

	"github.com/MobaiRuby/talk_it_api/cache"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/models"
	"github.com/MobaiRuby/talk_it_api/services/gobelieve"
	"github.com/MobaiRuby/talk_it_api/services/talk_it"
	"github.com/MobaiRuby/talk_it_api/services/talk_it_template"
	"github.com/henrylee2cn/faygo"
)

func (h *Base) Friends() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		cur_user_id, err := strconv.Atoi(ctx.Data(global.TALK_IT_HEADER_CURRENT_USER_ID).(string))
		if cur_user_id == 0 || err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeUserIdError,
				"msg":  "get current user failed",
			})
		}

		friends, err := talk_it.Friends(cur_user_id)
		if err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeFriendsError,
				"msg":  err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "",
			"data": friends,
		})
	})
}

func (h *Base) AddFriend() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		friend_req := &models.FriendRequest{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), friend_req); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeAddFriendError,
				"msg":  err.Error(),
			})
		}

		if friend_req.Content == "" {
			friend_req.Content = talk_it_template.GenTemplate(global.TEMPLATE_FRIEND_REQUEST, friend_req.FromUserId)
		}

		if _, err := gobelieve.SystemMsg(friend_req.ToUserId, friend_req.Content); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeAddFriendError,
				"msg":  err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, faygo.Map{
			"code": 0,
			"msg":  "add friend success",
		})
	})
}

func (h *Base) AcceptFriend() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		friend_accept := &models.FriendAccept{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), friend_accept); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeAcceptFriendError,
				"msg":  err.Error(),
			})
		}

		if friend, err := friend_accept.Handle(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeAcceptFriendError,
				"msg":  err.Error(),
			})
		} else {
			cache.GetFriendsCache().Store(friend.Id, friend)
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": 0,
				"msg":  "accept friend request success",
			})
		}
	})
}

func (h *Base) RemoveFriend() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		um := &models.UAndMe{}
		if err := ctxBindJson(ctx.Data(global.CTX_BODY).([]byte), um); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeRemoveFriendError,
				"msg":  err.Error(),
			})
		}

		if friend, err := um.Remove(); err != nil {
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": errorCodeRemoveFriendError,
				"msg":  err.Error(),
			})
		} else {
			cache.GetFriendsCache().Delete(friend.Id)
			return ctx.JSON(http.StatusOK, faygo.Map{
				"code": 0,
				"msg":  "remove friend success",
			})
		}
	})
}
