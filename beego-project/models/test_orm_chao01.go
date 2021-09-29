package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id     int    `json:"id" orm:"pk"`
	Name   string `json:"name" orm:"column(Name)"`
	Age    byte   `json:"age"`
	Adress string `json:"adress"`
	Job    string `json:"job"`
}

func init() {
	orm.RegisterModel(new(User))
}
func (u *User) TableName() string {
	return "beego_table01"
}
