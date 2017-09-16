package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Salt     string
	Password string `json:"password"`
}

func init() {
	orm.RegisterModelWithPrefix("devbus_", new(User))
}
