package facebook

import (
	"fmt"

	"net/http"

	"encoding/json"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

type FbToken struct {
	Id                      string `json:"id"`
	AccessToken             string `json:"access_token"`
	TokenRefreshIntervalSec int    `json:"token_refresh_interval_sec"`
}

func GetToken(auth_code string) (string, error) {
	path := helper.GetUrlPath(
		conf.TalkITConfig.Facebook.Domain,
		conf.TalkITConfig.Facebook.Url.Version,
		conf.TalkITConfig.Facebook.Url.AccessToken,
	)

	query := helper.GetUrlQuery(
		"grant_type=authorization_code",
		fmt.Sprintf("code=%s", auth_code),
		genSecret(),
	)

	url := helper.GetUrl(path, query)
	resp, err := utils.Hpool.Request(url, http.MethodGet, "", map[string]string{})
	if err != nil {
		return "", err
	}

	fb_token := &FbToken{}
	if err := json.Unmarshal([]byte(resp.(string)), fb_token); err != nil {
		return "", err
	}

	return fb_token.AccessToken, nil
}
