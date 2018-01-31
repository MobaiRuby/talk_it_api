package facebook

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

type Validation struct {
	Id    string `json:"id"`
	Phone struct {
		Number        string `json:"number"`
		CountryPrefix string `json:"country_prefix"`
		NationNumber  string `json:"nation_number"`
	} `json:"phone"`
	Application struct {
		Id string `json:"id"`
	} `json:"application"`
}

func ValidateToken(token string) (*Validation, error) {
	url := helper.GetUrl(helper.GetUrlPath(
		conf.TalkITConfig.Facebook.Domain,
		conf.TalkITConfig.Facebook.Url.Version,
		conf.TalkITConfig.Facebook.Url.Me,
	), helper.GetUrlQuery(
		fmt.Sprintf("access_token=%s", token),
	))

	resp, err := utils.Hpool.Request(url, http.MethodGet, "", map[string]string{})
	if err != nil {
		return nil, err
	}

	v := &Validation{}
	if err := json.Unmarshal([]byte(resp.(string)), v); err != nil {
		return nil, err
	}

	return v, nil
}
