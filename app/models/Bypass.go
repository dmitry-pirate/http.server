package models

import (
	"time"
	"vpn_api/app/store"
)

type Bypass struct {
	Id        int       `db:"id"`
	Host      string    `db:"host"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetAllBypass(s *store.Store) []Bypass {
	var bypass []Bypass
	if err := s.GetConnection().Select(&bypass, "select * from bypass"); err != nil {
		panic(err)
	}
	return bypass
}
