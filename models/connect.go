package models

import (
	"GinBlog/packages/ini"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	if ini.Datatype == "sqlite3" {
		DB, err = gorm.Open(sqlite.Open(ini.Datafile), &gorm.Config{})
		if err == nil {
			fmt.Println("数据连接成功")
		} else {
			fmt.Println("数据连接失败")
		}
	}
}
