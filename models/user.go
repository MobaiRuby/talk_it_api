package models

import (
	"time"

	"github.com/MobaiRuby/talk_it_api/db"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

type User struct {
	Id      int64
	Name    string    `xorm:"name" json:"name"`
	Tel     string    `xorm:"tel" json:"tel"`
	Token   string    `xorm:""`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
	Deleted time.Time `xorm:"deleted" json:"deleted"`
}

func init() {
	xorm.MustDB().Sync2(new(User))
}

func (u *User) IsNew() (bool, error) {
	if err := db.TalkITEngine.Native.Where("tel=?", u.Tel).Find(u); err != nil {
		return false, err
	}

	if u.Id > 0 {
		return false, nil
	}
	return true, nil
}

func (u *User) Insert() (*User, error) {
	if _, err := db.TalkITEngine.Native.InsertOne(u); err != nil {
		return nil, err
	}

	return u, nil
}
