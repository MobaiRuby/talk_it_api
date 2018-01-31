package gobelieve

import (
	"fmt"
	"net/http"

	"github.com/MobaiRuby/talk_it_api/conf"
	"github.com/MobaiRuby/talk_it_api/helper"
	"github.com/MobaiRuby/talk_it_api/utils"
)

/*
** 建群
** return {"group_id":"群组id(整型)"}
 */
func CreateGroup(master_id int, group_name string, is_super_group bool, member_ids ...int) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups)

	body, err := helper.Obj2Str(map[string]interface{}{
		"master":  master_id,
		"name":    group_name,
		"super":   is_super_group,
		"members": member_ids,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodPost, body, genAuthHeader())
}

// 设置群名
func SetGroupName(group_id int, group_name string) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups,
		fmt.Sprint(group_id))

	body, err := helper.Obj2Str(map[string]interface{}{
		"name": group_name,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodPatch, body, genAuthHeader())
}

// 解散群
func DismissGroup(group_id int) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups,
		fmt.Sprint(group_id))

	return utils.Hpool.Request(url, http.MethodDelete, "", genAuthHeader())
}

// 离开群
func LeaveGroup(group_id int, member_id int) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups,
		fmt.Sprint(group_id),
		conf.TalkITConfig.Gobelieve.Url.Members,
		fmt.Sprint(member_id))

	return utils.Hpool.Request(url, http.MethodDelete, "", genAuthHeader())
}

// 增加群成员
func AddGroupMembers(group_id int, member_ids ...int) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups,
		fmt.Sprint(group_id),
		conf.TalkITConfig.Gobelieve.Url.Members)

	body, err := helper.Obj2Str(map[string]interface{}{
		"uid": member_ids,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodPost, body, genAuthHeader())
}

// 移除群成员
func RemoveGroupMembers(group_id int, member_ids ...int) (interface{}, error) {
	url := helper.GetUrlPath(
		conf.TalkITConfig.Gobelieve.Domain,
		conf.TalkITConfig.Gobelieve.Url.Groups,
		fmt.Sprint(group_id),
		conf.TalkITConfig.Gobelieve.Url.Members)

	body, err := helper.Obj2Str(map[string]interface{}{
		"uid": member_ids,
	})
	if err != nil {
		return nil, err
	}

	return utils.Hpool.Request(url, http.MethodDelete, body, genAuthHeader())
}
