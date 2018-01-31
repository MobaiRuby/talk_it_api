package models

import "github.com/MobaiRuby/talk_it_api/db"

type Technoshpere struct {
	Id      int
	UserId  int    `xorm:"user_id" json:"user_id"`
	Content string `xorm:"content" json:"content"`
}

func Technoshperes() ([]*Technoshpere, error) {
	ts := []*Technoshpere{}
	err := db.TalkITEngine.Native.Where("").Find(&ts)
	return ts, err
}
