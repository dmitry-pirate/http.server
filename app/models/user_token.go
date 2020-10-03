package models

import "time"

type UserToken struct {
	Id         int       `db:"id"`
	Token      string    `db:"token"`
	LastActive time.Time `db:"last_active"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
