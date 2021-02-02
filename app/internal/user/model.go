package user

import (
	"database/sql"
)

type Users struct {
	Id                 int            `db:"id"`
	Name               sql.NullString `db:"name"`
	Email              string         `db:"email"`
	SubscriptionStatus string         `db:"subscription_status"`
}
