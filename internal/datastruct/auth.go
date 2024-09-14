package datastruct

import "github.com/golang-jwt/jwt/v5"

type AuthLoginData struct {
	Id       string `db:"uuid"`
	Username string `db:"username"`
	Password string `db:"password"`
	RoleCode string `db:"role_code"`
	RoleName string `db:"role_name"`
	IsActive bool   `db:"is_active"`
}

type AuthToken struct {
	AccessToken    *string
	AccessExpired  *jwt.NumericDate
	RefreshToken   *string
	RefreshExpired *jwt.NumericDate
}

type AuthMe struct {
	UUID     string `db:"uuid" json:"id"`
	Username string `db:"username" json:"username"`
	RoleCode string `db:"role_code" json:"role_code"`
	RoleName string `db:"role_name" json:"role_name"`
}
