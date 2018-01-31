package gobelieve

import (
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

/*
**  return {"token":"访问token"}
 */
func GetToken(user_id int64, user_name string) (string, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Auth,
		conf.TalkITConfig.Gobelieve.Url.Grant)

	body, err := helper.Obj2Str(map[string]interface{}{
		"uid":         user_id,
		"user_name":   user_name,
		"platform_id": getPlatformId(),
		"device_id":   "",
	})
	if err != nil {
		return "", err
	}

	resp, err := utils.Hpool.Request(url, http.MethodPost, body, genAuthHeader())
	if err != nil {
		return "", err
	}

	token := map[string]map[string]string{
		"data": {
			"token": "",
		},
	}
	if err := json.Unmarshal([]byte(resp.(string)), &token); err != nil {
		return "", err
	}

	return token["data"]["token"], nil
}

func SetUserName(user_id int, user_name string) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Users,
		fmt.Sprint(user_id))

	body, err := helper.Obj2Str(map[string]interface{}{
		"name": user_name,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodPatch, body, genAuthHeader())
}
