package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Cfg        *ini.File
	err        error
	App_mode   string
	Datafile   string
	Datatype   string
	Port       int
	JwtSecret  string
	ExpireTime int
)

func init() {
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("配置文件加载失败")
	}
	Datafile = Cfg.Section("database").Key("datafile").String()
	Datatype = Cfg.Section("database").Key("datatype").String()
	Port = Cfg.Section("server").Key("port").MustInt()
	App_mode = Cfg.Section("").Key("app_mode").String()
	//JwtSecret = Cfg.Section("jwt").Key("jwtSecret").String()
	ExpireTime = Cfg.Section("jwt").Key("expireTime").MustInt()

}
