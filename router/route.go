package router

import (
	"github.com/MobaiRuby/talk_it_api/filter"
	"github.com/MobaiRuby/talk_it_api/handler"
	"github.com/henrylee2cn/faygo"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	// 基础数据服务的handle
	baseHandle := new(handler.Base)

	frame.Filter(filter.Pre, filter.Validate).Route(
		frame.NewGET("/", handler.Index),
		frame.NewHEAD("/", handler.Index),
		frame.NewNamedGroup("4 client api", "/api",
			// auth
			frame.NewPOST("/login/tel", baseHandle.TelLogin()),
			frame.NewPOST("/login/wechat", baseHandle.WeChatLogin()),
			// 好友
			frame.NewGET("/friends", baseHandle.Friends()),
			frame.NewGET("/friends/:id", baseHandle.Friend()),
			frame.NewPOST("/friends", baseHandle.AddFriend()),
			frame.NewPOST("/friends/accept", baseHandle.AcceptFriend()),
			frame.NewDELETE("/friends/:id", baseHandle.RemoveFriend()),
			// 群组
			frame.NewGET("/groups", baseHandle.Groups()),
			frame.NewGET("/groups/:id", baseHandle.Group()),
			frame.NewPOST("/groups", baseHandle.CreateGroup()),
			frame.NewDELETE("/groups/:id", baseHandle.DismissGroup()),
			// 群成员
			frame.NewGET("/members", baseHandle.Members()),
			frame.NewGET("/members/:id", baseHandle.Member()),
			frame.NewPOST("/members", baseHandle.AddMember()),
			frame.NewDELETE("/members/:id", baseHandle.RemoveMember()),
			// 技术圈
			frame.NewGET("/technoshperes", baseHandle.Technoshperes()),
			frame.NewGET("/technoshpere/:id", baseHandle.Technoshpere()),
			frame.NewPOST("/technoshperes", baseHandle.AddTechnoshpere()),
			frame.NewDELETE("/technoshperes/:id", baseHandle.RemoveTechnoshpere()),
			// 个人中心
			frame.NewGET("/user_center", baseHandle.UserCenter()),
		),
	)
}
