package cache

import (
	"sync"

	"github.com/MobaiRuby/talk_it_api/models"
	"github.com/henrylee2cn/goutil"
)

var (
	wg                sync.WaitGroup
	friendsCache      goutil.Map
	technoshpereCache goutil.Map
	groupCache        goutil.Map
	memberCache       goutil.Map
)

func init() {
	friendsCache = goutil.AtomicMap()
	technoshpereCache = goutil.AtomicMap()
	groupCache = goutil.AtomicMap()
	memberCache = goutil.AtomicMap()
}

func LoadCache() {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		loadFriends(models.QueryFriends)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		loadTechnoshperes(models.Technoshperes)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		loadGroups(models.Groups)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		loadMembers(models.Members)
		wg.Done()
	}(&wg)

	wg.Wait()
}

func loadFriends(f func() ([]*models.Friend, error)) error {
	fs, err := f()
	if err != nil {
		return err
	}

	for _, f := range fs {
		friendsCache.Store(f.Id, f)
	}

	return nil
}

func loadTechnoshperes(f func() ([]*models.Technoshpere, error)) error {
	ts, err := f()
	if err != nil {
		return err
	}

	for _, t := range ts {
		technoshpereCache.Store(t.Id, t)
	}

	return nil
}

func loadGroups(f func() ([]*models.Group, error)) error {
	gs, err := f()
	if err != nil {
		return err
	}

	for _, g := range gs {
		technoshpereCache.Store(g.Id, g)
	}

	return nil
}

func loadMembers(f func() ([]*models.Member, error)) error {
	ms, err := f()
	if err != nil {
		return err
	}

	for _, m := range ms {
		technoshpereCache.Store(m.Id, m)
	}

	return nil
}

func GetFriendsCache() goutil.Map {
	return friendsCache
}

func GetTechnoshpereCache() goutil.Map {
	return technoshpereCache
}

func GetGroupCache() goutil.Map {
	return groupCache
}

func GetMemberCache() goutil.Map {
	return memberCache
}
