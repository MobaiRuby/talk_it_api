package helper

import (
	"encoding/json"
	"strings"

	"github.com/MobaiRuby/talk_it_api/global"

	"encoding/base64"

	"crypto/md5"
	"encoding/hex"
)

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

func Obj2Str(i interface{}) (string, error) {
	bs, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func GenAuthHeader() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + "token",
	}
}

func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func HexMd5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUrlPath(urls ...string) string {
	return strings.Join(urls, global.SEP_LEFT_SLASH)
}

func GetUrlQuery(queries ...string) string {
	return strings.Join(queries, global.SEP_URL_QUERY_JOIN)
}

func GetUrl(path, query string) string {
	return strings.Join([]string{path, query}, global.SEP_PATH_QUERY)
}
