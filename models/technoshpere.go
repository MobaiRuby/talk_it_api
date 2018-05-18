package models

import (
	"fmt"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type Technoshpere struct {
	Id      int
	UserId  int    `xorm:"user_id" json:"user_id"`
	Content string `xorm:"content" json:"content"`
}

func init() {
	xorm.MustDB().Sync(new(Technoshpere))
}

func Technoshperes() ([]*Technoshpere, error) {
	ts := []*Technoshpere{}
	err := db.TalkITEngine.Native.Where("").Find(&ts)
	return ts, err
}

func (t *Technoshpere) One() error {
	has, err := db.TalkITEngine.Native.Get(t)
	if !has || err != nil {
		return fmt.Errorf("get technoshpere info faild")
	}
	return nil
}

func (t *Technoshpere) Remove() error {
	_, err := db.TalkITEngine.Native.Delete(t)
	if err != nil {
		return fmt.Errorf("remove technoshpere info faild")
	}
	return nil
}
