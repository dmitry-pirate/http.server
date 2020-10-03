package models

import (
	"database/sql"
)

type Users struct {
	Id                 int            `db:"id"`
	Name               sql.NullString `db:"name"`
	Email              string         `db:"email"`
	SubscriptionStatus string         `db:"subscription_status"`
}

type UsersJson struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	SubscriptionStatus string `json:"subscription_status"`
}
