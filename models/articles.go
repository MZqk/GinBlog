package models

import (
	"GinBlog/packages/ini"
	"fmt"
	"time"
)

type Article struct {
	Id          int
	Title       string
	Context     string
	Created_on  time.Time `json:"created_on"`
	Created_by  string    `json:"created_by"`
	Modified_on time.Time `json:"modified_on"`
	Modified_by string    `json:"modified_by"`
	Delete_on   time.Time `json:"delete_on"`
	State       int       `json:"state"`
}

func init() {
	if ini.App_mode == "init" {
		_ = DB.AutoMigrate(&Article{})
		fmt.Println("articles表已初始化")
	}
}

func CreateArticle(context string) {
	DB.Create(&Article{Context: context})
}
