package models

type User struct {
	UserId    uint64 `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	RegisterTime uint64 `json:"register_time"`
}