package user

import (
	"github.com/basketforcode/http.server/app/internal/usertoken"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (rep *repository) FindByID(id int) (Users, error) {
	var usr Users
	if err := rep.db.Get(&usr, "select id, email, name, subscription_status from users where id = ? limit 1", id); err != nil {
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
