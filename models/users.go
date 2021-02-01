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

func Createuser(username, password string) {
	DB.Create(map[string]interface{}{"username": username, "password": password})
}

func Selectuser() (user []User) {
	DB.Debug().Select("username").Find(&user)
	return
}

func CheckAuth(username, password string) bool {
	var user User
	DB.Debug().Select("id").Where(map[string]interface{}{"username": username, "password": password}).Find(&user)
	if user.Id > 0 {
		return true
	}
	return false
}
