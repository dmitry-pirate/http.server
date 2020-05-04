package models

import (
	"database/sql"
	"time"
)

type Users struct {
	Id                  int            `db:"id"`
	Name                sql.NullString `db:"name"`
	Email               string         `db:"email"`
	EmailVerifiedAt     sql.NullTime   `db:"email_verified_at"`
	Password            sql.NullString `db:"password"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
	CanceledAt          sql.NullTime   `db:"canceled_at"`
	StartTrial          time.Time      `db:"start_trial"`
	EndTrial            time.Time      `db:"end_trial"`
	SubscriptionStatus  string         `db:"subscription_status"`
	FSAccount           sql.NullString `db:"fs_account"`
	ManageToken         string         `db:"manage_token"`
	Status              string         `db:"status"`
	UrlToken            string         `db:"url_token"`
	MarketingTag        sql.NullString `db:"marketing_tag"`
	CountryId           sql.NullInt64  `db:"country_id"`
	MailSubscribeStatus string         `db:"mail_subscribe_status"`
	MailUnsubscribedAt  sql.NullTime   `db:"mail_unsubscribe_at"`
	DevicesLimit        int            `db:"devices_limit"`
}

type UserToken struct {
	Id         int       `db:"id"`
	Token      string    `db:"token"`
	LastActive time.Time `db:"last_active"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
