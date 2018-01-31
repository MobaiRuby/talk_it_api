package easemob

import (
	"net/http"
	"strings"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

type Token struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func GetToken() (interface{}, error) {
	easemob_token := Token{
		GrantType:    conf.TalkITConfig.EaseMob.GrantType.ClientCredentials,
		ClientId:     conf.TalkITConfig.EaseMob.ClientID,
		ClientSecret: conf.TalkITConfig.EaseMob.ClientSecret,
	}

	body, err := helper.Obj2Str(easemob_token)
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Token}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, body, map[string]string{})
}

func PostUser(username, password string) (interface{}, error) {
	user := User{
		UserName: username,
		Password: password,
	}

	body, err := helper.Obj2Str(user)
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, body, helper.GenAuthHeader())
}

func ResetPassword(password string) (interface{}, error) {
	body, err := helper.Obj2Str(map[string]string{
		"newpassword": password,
	})
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Password}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPut, body, helper.GenAuthHeader())
}

func UserOnlineOrOffline(username string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		username,
		conf.TalkITConfig.EaseMob.Url.Status}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodGet, "", helper.GenAuthHeader())
}

func DisableUser(username string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		username,
		conf.TalkITConfig.EaseMob.Url.Deactivate}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, "", helper.GenAuthHeader())
}

func EnabaleUser(username string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		username,
		conf.TalkITConfig.EaseMob.Url.Activate}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, "", helper.GenAuthHeader())
}

func ForceUserOffline(username string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		username,
		conf.TalkITConfig.EaseMob.Url.Activate}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, "", helper.GenAuthHeader())
}
