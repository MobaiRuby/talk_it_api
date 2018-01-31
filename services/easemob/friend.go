package easemob

import (
	"strings"

	"net/http"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

func MakeFriends(from, to string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Contacts,
		conf.TalkITConfig.EaseMob.Url.Users,
		to}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, "", helper.GenAuthHeader())
}

func BreakFriends(from, to string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Contacts,
		conf.TalkITConfig.EaseMob.Url.Users,
		to}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodDelete, "", helper.GenAuthHeader())
}

func GetFriends(from string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Contacts,
		conf.TalkITConfig.EaseMob.Url.Users}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodGet, "", helper.GenAuthHeader())
}

func GetBlocks(from string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Blocks,
		conf.TalkITConfig.EaseMob.Url.Users}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodGet, "", helper.GenAuthHeader())
}

func ToBlocks(from, to string) (interface{}, error) {
	body, err := helper.Obj2Str(map[string][]string{
		"usernames": {to},
	})
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Blocks,
		conf.TalkITConfig.EaseMob.Url.Users}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodPost, body, helper.GenAuthHeader())
}

func LeaveBlocks(from, to string) (interface{}, error) {
	url := strings.Join([]string{conf.TalkITConfig.EaseMob.Domain,
		conf.TalkITConfig.EaseMob.OrgName,
		conf.TalkITConfig.EaseMob.AppName,
		conf.TalkITConfig.EaseMob.Url.Users,
		from,
		conf.TalkITConfig.EaseMob.Url.Blocks,
		conf.TalkITConfig.EaseMob.Url.Users,
		to}, global.SEP_LEFT_SLASH)

	return utils.Hpool.Request(url, http.MethodDelete, "", helper.GenAuthHeader())
}
