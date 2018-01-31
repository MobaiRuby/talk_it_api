package main

import (
	"github.com/MobaiRuby/talk_it_api/cache"
	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/router"
	"github.com/henrylee2cn/faygo"
)

func init() {
	conf.LoadConfig()
	cache.LoadCache()
}

func main() {
	router.Route(faygo.New("talk_it_api"))
	faygo.Run()
}
