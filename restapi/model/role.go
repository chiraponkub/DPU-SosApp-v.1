package model

type AddRole struct {
	Name string `json:"name" validate:"required"`
}

type GetRoleList struct {
	Name string `json:"name"`
}
