package usertoken

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{db: db}
}

//Return true if auth token is exists in db
func (rep *repository) Get(token string) (UserToken, error) {
	var userToken UserToken

	if err := rep.db.Get(&userToken, "select * from user_token where token = ?", token); err != nil {
		return userToken, err
	}

	return userToken, nil
}
