package models

type User struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Salt       string
	Password   string `json:"password"`
	IsActivate bool   `json:"is_activate"`
}

func (User) TableName() string {
	return "devbus_user"
}
