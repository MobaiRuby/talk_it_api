package models

import (
	"fmt"
	"time"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type Group struct {
	Id           int       `xorm:"id" json:"id"`
	Name         string    `xorm:"name" json:"name"`
	Announcement string    `xorm:"announcement" json:"announcement"`
	UserId       int       `xorm:"user_id" json:"user_id"` // 群主
	Created      time.Time `xorm:"created" json:"created"`
	Updated      time.Time `xorm:"updated" json:"updated"`
	Deleted      time.Time `xorm:"deleted" json:"deleted"`
}

func init() {
	xorm.MustDB().Sync(new(Group))
}

func Groups() ([]*Group, error) {
	groups := []*Group{}
	err := db.TalkITEngine.Native.Where("").Find(&groups)
	return groups, err
}

func (g *Group) Create() error {
	_, err := db.TalkITEngine.Native.InsertOne(g)
	return err
}

func (g *Group) Dismiss() error {
	return nil
}

func (g *Group) One() error {
	has, err := db.TalkITEngine.Native.Get(g)
	if !has || err != nil {
		return fmt.Errorf("get friend info faild")
	}
	return nil
}
