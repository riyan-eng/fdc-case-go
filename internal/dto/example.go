package dto

type ExampleCreate struct {
	Name   string `json:"name" valid:"required;min:3" minLength:"3" validate:"required"`
	Detail string `json:"detail"`
}

type ExamplePatch struct {
	Name   string `json:"name" valid:"min:3" minLength:"3"`
	Detail string `json:"detail"`
}

type ExamplePut struct {
	Name   string `json:"name" valid:"required;min:3" minLength:"3" validate:"required"`
	Detail string `json:"detail"`
}
