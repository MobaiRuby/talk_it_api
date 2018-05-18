package models

import (
	"fmt"
	"time"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type Member struct {
	Id         int
	MasterId   int       `xorm:"master_id" json:"master_id"`
	MasterName string    `xorm:"master_name" json:"master_name"`
	UserId     int       `xorm:"user_id" json:"user_id"`
	UserName   string    `xorm:"user_name" json:"user_name"`
	GroupId    int       `xorm:"group_id" json:"group_id"`
	GroupName  string    `xorm:"group_name" json:"group_name"`
	Created    time.Time `xorm:"created" json:"created"`
	Updated    time.Time `xorm:"updated" json:"updated"`
	Deleted    time.Time `xorm:"deleted" json:"deleted"`
}

func init() {
	xorm.MustDB().Sync(new(Member))
}

func Members() ([]*Member, error) {
	members := []*Member{}
	err := db.TalkITEngine.Native.Where("").Find(&members)
	return members, err
}

func (m *Member) Add() error {
	_, err := db.TalkITEngine.Native.InsertOne(m)
	return err
}

func (m *Member) Remove() error {
	_, err := db.TalkITEngine.Native.Delete(m)
	return err
}

func (m *Member) Leave() error {
	_, err := db.TalkITEngine.Native.Delete(m)
	return err
}

func (m *Member) One() error {
	has, err := db.TalkITEngine.Native.Get(m)
	if !has || err != nil {
		return fmt.Errorf("get friend info faild")
	}
	return nil
}
