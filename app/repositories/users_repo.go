package repositories

import (
	"github.com/dmitry-pirate/http.server/app/models"
	"github.com/dmitry-pirate/http.server/app/services/store"
)

type usersRepo struct {
	store *store.Store
}

func NewUserRepo(store *store.Store) *usersRepo {
	return &usersRepo{store: store}
}

//Return user by auth token from header
func (rep *usersRepo) GetByID(id int) (models.Users, error) {
	var usr models.Users
	if err := rep.store.SlaveConnection().Get(&usr, "select * from users where id = ? limit 1", id); err != nil {
		return usr, err
	}
	return usr, nil
}

func (rep *usersRepo) GetFormattedInfo(ut models.UserToken) (*models.UsersJson, error) {
	var usrJSON models.UsersJson
	usr, err := rep.GetByID(ut.Id)
	if err != nil {
		return nil, err
	}

	usrJSON = models.UsersJson{
		ID:                 usr.Id,
		Email:              usr.Email.String,
		Name:               usr.Name.String,
		SubscriptionStatus: usr.SubscriptionStatus,
	}
	return &usrJSON, nil
}
