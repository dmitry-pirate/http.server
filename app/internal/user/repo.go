package user

import (
	"github.com/basketforcode/http.server/app/internal/usertoken"
	"github.com/basketforcode/http.server/app/services/store"
)

type repository struct {
	store *store.Store
}

func NewRepo(store *store.Store) *repository {
	return &repository{store: store}
}

//Return user by auth token from header
func (rep *repository) FindByID(id int) (Users, error) {
	var usr Users
	if err := rep.store.SlaveConnection().Get(&usr, "select id, email, name, subscription_status from users where id = ? limit 1", id); err != nil {
		return usr, err
	}
	return usr, nil
}

func (rep *repository) FormatInfo(ut usertoken.UserToken) (*ResponseJson, error) {
	var usrJSON ResponseJson
	usr, err := rep.FindByID(ut.Id)
	if err != nil {
		return nil, err
	}

	usrJSON = ResponseJson{
		ID:                 usr.Id,
		Email:              usr.Email,
		Name:               usr.Name.String,
		SubscriptionStatus: usr.SubscriptionStatus,
	}
	return &usrJSON, nil
}
