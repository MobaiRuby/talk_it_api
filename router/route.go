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
			frame.NewPOST("/friends/add", baseHandle.AddFriend()),
			frame.NewPOST("/friends/accept", baseHandle.AcceptFriend()),
			frame.NewPOST("/friends/remove", baseHandle.RemoveFriend()),
			// 群组
			frame.NewGET("/groups", baseHandle.Groups()),
			frame.NewPOST("/groups/create", baseHandle.CreateGroup()),
			frame.NewPOST("/groups/dismiss", baseHandle.DismissGroup()),
			// 群成员
			frame.NewGET("/groups/members/:group_id", baseHandle.Members()),
			frame.NewPOST("/groups/members/:group_id/add", baseHandle.AddMember()),
			frame.NewPOST("/groups/members/:group_id/remove", baseHandle.RemoveMember()),
			frame.NewPOST("/groups/members/:group_id/leave", baseHandle.LeaveGroup()),
			// 技术圈
			frame.NewGET("/technoshperes", baseHandle.Technoshperes()),
			// 个人中心
			frame.NewGET("/user_center", baseHandle.UserCenter()),
		),
	)
}
