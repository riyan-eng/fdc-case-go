package model

import "database/sql"

type Example struct {
	Id     string         `db:"id"`
	Name   sql.NullString `db:"name"`
	Detail sql.NullString `db:"detail"`
}
