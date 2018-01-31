package easemob

import "fmt"

type ChatGroup struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Public            bool              `json:"public"`
	MembersOnly       bool              `json:"membersonly"`
	AllowInvites      bool              `json:"allowinvites"`
	MaxUsers          int               `json:"maxusers"`
	AffiliationsCount int               `json:"affiliations_count"`
	Affiliations      map[string]string `json:"affiliations"`
	Owner             string            `json:"owner"`
	Member            string            `json:"member"`
	InviteNeedConfirm bool              `json:"invite_need_confirm"`
}

func GetAllChatGroups() (interface{}, error) {
	return nil, fmt.Errorf("")
}
