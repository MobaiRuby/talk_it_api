package talk_it

import (
	"fmt"

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

func Friend(id int) (*models.Friend, error) {
	v, ok := cache.GetFriendsCache().Load(id)
	if v != nil && ok {
		return v.(*models.Friend), nil
	}

	return nil, fmt.Errorf("get friend info faild")
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

func Member(id int) (*models.Member, error) {
	v, ok := cache.GetMemberCache().Load(id)
	if v != nil && ok {
		return v.(*models.Member), nil
	}

	return nil, fmt.Errorf("get member info faild")
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

func Group(id int) (*models.Group, error) {
	v, ok := cache.GetGroupCache().Load(id)
	if v != nil && ok {
		return v.(*models.Group), nil
	}

	return nil, fmt.Errorf("get group info faild")
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

func Technoshpere(id int) (*models.Technoshpere, error) {
	v, ok := cache.GetTechnoshpereCache().Load(id)
	if v != nil && ok {
		return v.(*models.Technoshpere), nil
	}

	return nil, fmt.Errorf("get technoshpere info faild")
}
