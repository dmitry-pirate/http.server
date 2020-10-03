package models

import (
	"database/sql"
	"time"
)

type Users struct {
	Id                 int            `db:"id"`
	Name               sql.NullString `db:"name"`
	Email              sql.NullString `db:"email"`
	RememberToken      sql.NullString `db:"remember_token"`
	EmailVerifiedAt    sql.NullTime   `db:"email_verified_at"`
	Password           sql.NullString `db:"password"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
	CanceledAt         sql.NullTime   `db:"canceled_at"`
	StartTrial         sql.NullTime   `db:"start_trial"`
	EndTrial           sql.NullTime   `db:"end_trial"`
	SubscriptionStatus string         `db:"subscription_status"`
	FSAccount          sql.NullString `db:"fs_account"`
	ManageToken        string         `db:"manage_token"`
	Status             string         `db:"status"`
	UrlToken           string         `db:"url_token"`
	MarketingTag       sql.NullString `db:"marketing_tag"`
	CountryId          sql.NullInt64  `db:"country_id"`
	RegThrough         sql.NullString `db:"reg_through"`
	DevicesLimit       int            `db:"devices_limit"`
}

type UsersJson struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	SubscriptionStatus string `json:"subscription_status"`
}
