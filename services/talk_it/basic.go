package talk_it

import (
	"github.com/MobaiRuby/talk_it_api/cache"
	"github.com/MobaiRuby/talk_it_api/models"
)

func Friends(user_id int) (map[int]*models.Friend, error) {
	friends := map[int]*models.Friend{}

	cache.GetFriendsCache().Range(func(key, value interface{}) bool {
		if value != nil && (value.(models.Friend).FromUserId == user_id || value.(models.Friend).ToUserId == user_id) {
			friends[value.(models.Friend).Id] = value.(*models.Friend)
		}
		return true
	})

	return friends, nil
}

func Members(group_id int) (map[int]*models.Member, error) {
	members := map[int]*models.Member{}

	cache.GetMemberCache().Range(func(key, value interface{}) bool {
		if value != nil && value.(models.Member).GroupId == group_id {
			members[value.(models.Member).Id] = value.(*models.Member)
		}
		return true
	})

	return members, nil
}

func Groups(user_id int) (map[int]*models.Group, error) {
	groups := map[int]*models.Group{}

	cache.GetGroupCache().Range(func(key, value interface{}) bool {
		if value != nil && value.(models.Group).UserId == user_id {
			groups[value.(models.Group).Id] = value.(*models.Group)
		}
		return true
	})

	return groups, nil
}

func Technoshperes(user_id int) (map[int]*models.Technoshpere, error) {
	technoshperes := map[int]*models.Technoshpere{}

	cache.GetTechnoshpereCache().Range(func(key, value interface{}) bool {
		if value != nil && value.(models.Technoshpere).UserId == user_id {
			technoshperes[value.(models.Technoshpere).Id] = value.(*models.Technoshpere)
		}
		return true
	})

	return technoshperes, nil
}
