package db

import (
	"github.com/MobaiRuby/talk_it_api/global"
	native "github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type Engine struct {
	DbName string
	Native *native.Engine
}

var (
	TalkITEngine *Engine
)

func init() {
	TalkITEngine = &Engine{DbName: global.DB_NAME, Native: GetEngine(global.DB_NAME)}
}

func GetEngine(db_name string) *native.Engine {
	engine := xorm.MustDB(db_name)
	engine.ShowSQL(true)
	return engine
}
