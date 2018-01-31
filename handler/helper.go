package handler

import (
	"encoding/json"
)

const (
	errorCodeParamsError = iota - 1
	errorCodeFbTokenError
	errorCodeFbValidationError
	errorCodeRegisterUserError

	errorCodeUserIdError
	errorCodeFriendsError
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
)

const (
	fbAuthCode  = "fb_auth_code"
	fbAuthToken = "fb_auth_token"
)

func ctxBindJson(bs interface{}, target interface{}) error {
	return json.Unmarshal(bs.([]byte), target)
}
