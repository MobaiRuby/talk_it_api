package gobelieve

import (
	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
)

var (
	authHeader    map[string]string
	authorization = "Authorization"
	contentType   = "Content-Type"
)

func init() {
	authHeader = map[string]string{}
}

func getPlatformId() string {
	return "1"
}

func genAuthHeader() map[string]string {
	if _, ok := authHeader[authorization]; !ok {
		authHeader[authorization] = "Basic " + helper.Base64Encode(conf.TalkITConfig.Gobelieve.AppId+":"+helper.HexMd5(conf.TalkITConfig.Gobelieve.AppSecret))
	}

	if _, ok := authHeader[contentType]; !ok {
		authHeader[contentType] = "application/json; charset=UTF-8"
	}

	return authHeader
}
