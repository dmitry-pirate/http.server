package usertoken

import (
	"github.com/basketforcode/http.server/app/services/store"
)

type repository struct {
	store *store.Store
}

func NewRepo(store *store.Store) *repository {
	return &repository{store: store}
}

//Return true if auth token is exists in db
func (rep *repository) Get(token string) (UserToken, error) {
	var userToken UserToken

	if err := rep.store.SlaveConnection().Get(&userToken, "select * from user_token where token = ?", token); err != nil {
		return userToken, err
	}

	return userToken, nil
}
