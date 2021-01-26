package models

import (
	"GinBlog/packages/ini"
	"fmt"
)

type User struct {
	Id       int
	Username string
	Password string
	State    int
}

func init() {
	if ini.App_mode == "init" {
		_ = DB.AutoMigrate(&User{})
		fmt.Println("users表已初始化")
	}
}

func Createuser() {
	DB.Create(&User{Username: "mz", Password: "1q2w3e"})
}

func Selectuser() (user []User) {
	DB.Model(&User{}).Find(&user)
	return
}
