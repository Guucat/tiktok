package model

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
