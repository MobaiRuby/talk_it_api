package conf

import (
	"os"

	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/spf13/viper"
)

// traffic_unity_center 全局配置
type Config struct {
	EaseMob struct { // 环信配置
		Domain       string `mapstructure:"domain"`
		OrgName      string `mapstructure:"org_name"`
		AppName      string `mapstructure:"app_name"`
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
		Token        string `mapstructure:"client_secret"`
		GrantType    struct {
			ClientCredentials string `json:"client_credentials"`
		}
		Url struct {
			Token      string `mapstructure:"token"`
			Users      string `mapstructure:"users"`
			Password   string `mapstructure:"password"`
			Contacts   string `mapstructure:"contacts"`
			Blocks     string `mapstructure:"blocks"`
			Status     string `mapstructure:"status"`
			Deactivate string `mapstructure:"deactivate"`
			Activate   string `mapstructure:"activate"`
			Disconnect string `mapstructure:"disconnect"`
			ChatGroup  string `mapstructure:"chatgroup"`
		}
	} `mapstructure:"easemob"`
	Gobelieve struct {
		Domain    string `mapstructure:"domain"`
		AppId     string `mapstructure:"app_id"`
		AppKey    string `mapstructure:"app_key"`
		AppSecret string `mapstructure:"app_secret"`
		Url       struct {
			Auth    string `mapstructure:"auth"`
			Grant   string `mapstructure:"grant"`
			Users   string `mapstructure:"users"`
			Groups  string `mapstructure:"groups"`
			Members string `mapstructure:"members"`
			Message string `mapstructure:"message"`
			System  string `mapstructure:"system"`
		} `mapstructure:"url"`
	}
	Facebook struct {
		Domain      string `mapstructure:"domain"`
		AppId       string `mapstructure:"app_id"`
		AppSecret   string `mapstructure:"app_secret"`
		ClientToken string `mapstructure:"client_token"`
		Url         struct {
			Version     string `mapstructure:"version"`
			AccessToken string `mapstructure:"access_token"`
			Me          string `mapstructure:"me"`
		} `mapstructure:"url"`
	} `mapstructure:"facebook"`
}

var (
	TalkITConfig Config
)

func LoadConfig() {
	filename := os.Getenv(global.ENV_CONFIG_FILE_PATH)
	if filename == "" {
		filename = global.DEFAULT_CONFIG_FILE_PATH
	}

	viper.SetConfigType(global.DEFAULT_CONFIG_FILE_TYPE)
	viper.SetConfigFile(filename)

	helper.Must(nil, viper.ReadInConfig())
	helper.Must(nil, viper.Unmarshal(&TalkITConfig))
}
