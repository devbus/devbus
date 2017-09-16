package models

type Project struct {
	ID          int32
	Name        string `valid:"Required;Max(256)"`
	Description string
	DescType    TextType
}

func (Project) TableName() string {
	return "project"
}