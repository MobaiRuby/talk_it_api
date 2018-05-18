package handler

import (
	"encoding/json"

	"fmt"

	"strconv"

	"github.com/henrylee2cn/faygo"
)

const (
	errorCodeParamsError = iota - 1
	errorCodeFbTokenError
	errorCodeFbValidationError
	errorCodeRegisterUserError

	errorCodeUserIdError
	errorCodeFriendsError
	errorCodeFriendError
	errorCodeAddFriendError
	errorCodeAcceptFriendError
	errorCodeRemoveFriendError

	errorCodeGroupsError
	errorCodeGroupCreateError
	errorCodeGroupLeaveError
	errorCodeGroupDismissError

	errorCodeMembersError
	errorCodeMembersAddError
	errorCodeMembersRemoveError

	errorCodeTechnoshperesError
	errorCodeTechnoshpereError
	errorCodeTechnoshpereAddError
	errorCodeTechnoshpereRemoveError
)

const (
	fbAuthCode  = "fb_auth_code"
	fbAuthToken = "fb_auth_token"
)

const (
	pathIdKey string = ":id"
)

func ctxBindJson(bs interface{}, target interface{}) error {
	return json.Unmarshal(bs.([]byte), target)
}

func pathId(ctx *faygo.Context) (int, error) {
	path_id := ctx.PathParam(pathIdKey)
	if path_id == "" {
		return 0, fmt.Errorf("get path :id faild")
	}
	return strconv.Atoi(path_id)
}
