package repositories

import (
	"github.com/basketforcode/http.server/app/models"
	"github.com/basketforcode/http.server/app/services/store"
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
	if err := rep.store.SlaveConnection().Get(&usr, "select id, email, name, subscription_status from users where id = ? limit 1", id); err != nil {
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
		Email:              usr.Email,
		Name:               usr.Name.String,
		SubscriptionStatus: usr.SubscriptionStatus,
	}
	return &usrJSON, nil
}
