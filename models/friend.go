package models

import (
	"time"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/MobaiRuby/talk_it_api/global"
	"github.com/MobaiRuby/talk_it_api/services/gobelieve"
	"github.com/MobaiRuby/talk_it_api/services/talk_it_template"
)

type Friend struct {
	Id int `xorm:"id" json:"id"`
	UAndMe
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
	Deleted time.Time `xorm:"deleted" json:"deleted"`
}

// 好友请求
type FriendRequest struct {
	UAndMe
	Content string    `xorm:"content" json:"content"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
	Deleted time.Time `xorm:"deleted" json:"deleted"`
}

// 好友反馈
type FriendAccept struct {
	UAndMe
	Content string    `xorm:"content" json:"content"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
	Deleted time.Time `xorm:"deleted" json:"deleted"`
}

type UAndMe struct {
	FromUserId int `xorm:"from_user_id" json:"from_user_id"`
	ToUserId   int `xorm:"to_user_id" json:"to_user_id"`
}

func QueryFriends() ([]*Friend, error) {
	friends := []*Friend{}
	err := db.TalkITEngine.Native.Where("").Find(&friends)
	return []*Friend{}, err
}

func (fa *FriendAccept) Handle() (*Friend, error) {
	session := db.TalkITEngine.Native.NewSession()
	if err := session.Begin(); err != nil {
		return nil, err
	}
	friend := &Friend{
		UAndMe: UAndMe{
			FromUserId: fa.ToUserId,
			ToUserId:   fa.FromUserId,
		},
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
		UAndMe: UAndMe{
			FromUserId: um.FromUserId,
			ToUserId:   um.ToUserId,
		},
	}
	_, err := db.TalkITEngine.Native.Delete(remove_friend)
	if err != nil {
		return nil, err
	}
	return remove_friend, nil
}
