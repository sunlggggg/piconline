package entity

import "log"

type User struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisterTime uint64 `json:"register_time"`
}

func (user *User) Init(id uint64, name string, password string, email string, registerTime uint64) *User {
	if user == nil {
		log.Fatal("user 为 nil")
	}
	user.Name = name
	user.Email = email
	user.Password = password
	user.RegisterTime = registerTime
	user.Id = id
	return user
}
