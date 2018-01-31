package talk_it_template

import (
	"fmt"

	"github.com/MobaiRuby/talk_it_api/global"
)

func GenTemplate(template_type string, from interface{}) string {
	switch template_type {
	case global.TEMPLATE_FRIEND_REQUEST:
		return fmt.Sprintf("%s 添加你为好友", from)
	case global.TEMPLATE_FRIEND_ACCEPT:
		return fmt.Sprintf("%s 同意你的好友申请", from)
	default:
		return ""
	}
}
