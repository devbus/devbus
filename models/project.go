package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id          int32
	Name        string `valid:"Required;Max(256)"`
	Description string
	DescType    TextType
}

func init() {
	orm.RegisterModel(new(Project))
}
