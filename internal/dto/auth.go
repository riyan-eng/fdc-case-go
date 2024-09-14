package dto

type AuthLogin struct {
	Username string `json:"username" valid:"required" validate:"required"`
	Password string `json:"password" valid:"required" format:"password" validate:"required"`
}

type AuthRefresh struct {
	Token string `json:"token" valid:"required" validate:"required"`
}

type AuthRegister struct {
	Email    string `json:"email" valid:"required;email" format:"email" validate:"required"`
	Username string `json:"username" valid:"required;min:5" minLength:"5" validate:"required"`
	Password string `json:"password" valid:"required;min:8" minLength:"8" format:"password" validate:"required"`
}
