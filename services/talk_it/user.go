package talk_it

import (
	"github.com/MobaiRuby/talk_it_api/models"
	"github.com/MobaiRuby/talk_it_api/services/gobelieve"
)

func Register2TalkIt(tel string) (*models.User, error) {
	user := &models.User{Tel: tel}
	if is_new, err := user.IsNew(); err != nil {
		return nil, err
	} else if is_new {
		user.Name = "Tel:" + tel
		user, err := user.Insert()
		if err != nil {
			return nil, err
		}
		user.Token, err = gobelieve.GetToken(user.Id, user.Name)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
