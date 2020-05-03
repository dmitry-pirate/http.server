package models

import "time"

type Countries struct {
	Id          int       `db:"id"`
	CountryCode string    `db:"country_code"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
