package repositories

import (
	"github.com/dmitry-pirate/http.server/app/models"
	"github.com/dmitry-pirate/http.server/app/services/store"
)

type userTokenRepo struct {
	store *store.Store
}

func NewUserTokenRepo(store *store.Store) *userTokenRepo {
	return &userTokenRepo{store: store}
}

//Return true if auth token is exists in db
func (rep *userTokenRepo) Get(token string) (models.UserToken, error) {
	var userToken models.UserToken

	if err := rep.store.SlaveConnection().Get(&userToken, "select * from user_token where token = ?", token); err != nil {
		return userToken, err
	}

	return userToken, nil
}
