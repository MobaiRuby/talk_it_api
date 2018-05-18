package models

import (
	"time"

	"fmt"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/services/gobelieve"
	"github.com/MobaiRuby/talk_it_api/services/talk_it_template"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type Friend struct {
	Id           int       `xorm:"id" json:"id"`
	FromUserId   int       `xorm:"from_user_id" json:"from_user_id"`
	FromUserName string    `xorm:"from_user_name" json:"from_user_name"`
	ToUserId     int       `xorm:"to_user_id" json:"to_user_id"`
	ToUserName   string    `xorm:"to_user_name" json:"to_user_name"`
	Created      time.Time `xorm:"created" json:"created"`
	Updated      time.Time `xorm:"updated" json:"updated"`
	Deleted      time.Time `xorm:"deleted" json:"deleted"`
}

// 好友请求
type FriendRequest struct {
	FromUserId   int       `xorm:"from_user_id" json:"from_user_id"`
	FromUserName string    `xorm:"from_user_name" json:"from_user_name"`
	ToUserId     int       `xorm:"to_user_id" json:"to_user_id"`
	ToUserName   string    `xorm:"to_user_name" json:"to_user_name"`
	Content      string    `xorm:"content" json:"content"`
	Created      time.Time `xorm:"created" json:"created"`
	Updated      time.Time `xorm:"updated" json:"updated"`
	Deleted      time.Time `xorm:"deleted" json:"deleted"`
}

// 好友反馈
type FriendAccept struct {
	FromUserId   int       `xorm:"from_user_id" json:"from_user_id"`
	FromUserName string    `xorm:"from_user_name" json:"from_user_name"`
	ToUserId     int       `xorm:"to_user_id" json:"to_user_id"`
	ToUserName   string    `xorm:"to_user_name" json:"to_user_name"`
	Content      string    `xorm:"content" json:"content"`
	Created      time.Time `xorm:"created" json:"created"`
	Updated      time.Time `xorm:"updated" json:"updated"`
	Deleted      time.Time `xorm:"deleted" json:"deleted"`
}

type UAndMe struct {
	FromUserId   int    `xorm:"from_user_id" json:"from_user_id"`
	FromUserName string `xorm:"from_user_name" json:"from_user_name"`
	ToUserId     int    `xorm:"to_user_id" json:"to_user_id"`
	ToUserName   string `xorm:"to_user_name" json:"to_user_name"`
}

func init() {
	xorm.MustDB().Sync(new(Friend))
	xorm.MustDB().Sync(new(FriendAccept))
	xorm.MustDB().Sync(new(FriendRequest))
}

func QueryFriends() ([]*Friend, error) {
	friends := []*Friend{}
	err := db.TalkITEngine.Native.Where("").Find(&friends)
	return []*Friend{}, err
}

func (f *Friend) One() error {
	has, err := db.TalkITEngine.Native.Get(f)
	if !has || err != nil {
		return fmt.Errorf("get friend info faild")
	}
	return nil
}

func (fa *FriendAccept) Handle() (*Friend, error) {
	session := db.TalkITEngine.Native.NewSession()
	if err := session.Begin(); err != nil {
		return nil, err
	}
	friend := &Friend{
		FromUserId: fa.ToUserId,
		ToUserId:   fa.FromUserId,
	}
	if _, err := db.TalkITEngine.Native.InsertOne(friend); err != nil {
		session.Rollback()
		return nil, err
	}

	if _, err := gobelieve.SystemMsg(fa.ToUserId, talk_it_template.GenTemplate(global.TEMPLATE_FRIEND_ACCEPT, fa.FromUserId)); err != nil {
		session.Rollback()
		return nil, err
	}

	session.Commit()

	return friend, nil
}

func (um *UAndMe) Remove() (*Friend, error) {
	remove_friend := &Friend{
		FromUserId: um.FromUserId,
		ToUserId:   um.ToUserId,
	}
	_, err := db.TalkITEngine.Native.Delete(remove_friend)
	if err != nil {
		return nil, err
	}
	return remove_friend, nil
}
