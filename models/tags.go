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

func Createtag() {
	DB.Create(&Tag{Name: "test2", Created_by: "mz", Created_on: time.Now(), State: 1})
}

func Selecttag() (tag []Tag) {
	DB.Model(&Tag{}).Find(&tag)
	return
}

type tagname struct {
	Name string `json:"name"`
}

func Gettag(id int) (tag []tagname) {
	DB.Model(&Tag{}).Select("name").Where("id = ?", id).First(&tag)
	fmt.Println(tag)
	return
}
