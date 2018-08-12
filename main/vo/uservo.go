package vo

type User struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RegisterTime uint64 `json:"register_time"`
}
