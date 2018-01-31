package gobelieve

import (
	"net/http"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

func SystemMsg(receiver_id int, content string) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Message,
		conf.TalkITConfig.Gobelieve.Url.System)

	body, err := helper.Obj2Str(map[string]interface{}{
		"receiver": receiver_id,
		"content":  content,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodPost, body, genAuthHeader())
}
