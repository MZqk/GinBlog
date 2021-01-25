package models

import (
	"GinBlog/packages/ini"
	"fmt"
	"time"
)

type Tag struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Created_on  time.Time `json:"created_on"`
	Created_by  string    `json:"created_by"`
	Modified_on time.Time `json:"modified_on"`
	Modified_by string    `json:"modified_by"`
	Delete_on   time.Time `json:"delete_on"`
	State       int       `json:"state"`
}

func init() {
	if ini.App_mode == "init" {
		_ = DB.AutoMigrate(&Tag{})
		fmt.Println("tags表已初始化")
	}
}

func Createorm() {
	DB.Create(&Tag{Name: "test2", Created_by: "mz", Created_on: time.Now(), State: 1})
}

func Selectorm() (tag []Tag) {
	DB.Model(&Tag{}).Find(&tag)
	return
}
