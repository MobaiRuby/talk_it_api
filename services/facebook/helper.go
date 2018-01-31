package facebook

import (
	"strings"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/global"
)

var (
	prefix = "AA"
)

func genSecret() string {
	return strings.Join([]string{prefix, conf.TalkITConfig.Facebook.AppId, conf.TalkITConfig.Facebook.AppSecret}, global.SEP_VERTICAL_LINE)
}
