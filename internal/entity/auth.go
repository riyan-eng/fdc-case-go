package entity

type AuthRegister struct {
	UserId   *string
	Email    *string
	UserName *string
	Password *string
	RoleCode *string
}

type AuthLogin struct {
	Username *string
	Password *string
}

type AuthRefresh struct {
	Token *string
}

type AuthResetPassword struct {
	Token    *string
	Password *string
}

type AuthResetPasswordToken struct {
	Email *string
}

type AuthResetPasswordTokenValidate struct {
	Token *string
}

type AuthLogout struct {
	UserId *string
}

type AuthMe struct {
	UserId *string
}
