package model

import (
	"database/sql"
)

type User struct {
	Id        string         `db:"id"`
	Username  sql.NullString `db:"username"`
	Password  sql.NullString `db:"password"`
	IsActive  sql.NullBool   `db:"is_active"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

type UserData struct {
	Id       string         `db:"id"`
	UserId   sql.NullString `db:"user_id"`
	RoleCode sql.NullString `db:"role_code"`
}
