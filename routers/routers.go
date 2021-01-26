package routers

import (
	"GinBlog/packages/ini"
	v1 "GinBlog/routers/api/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"strconv"
)

func Routers() {
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// @title 测试
	// @version 0.0.1
	// @description  测试
	// @BasePath /api/v1/
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tag", v1.Createtag)
		apiv1.GET("/tag/:id", v1.Gettag)
		apiv1.GET("/tags", v1.Selecttag)
		apiv1.POST("/user", v1.Createuser)
		apiv1.GET("/users", v1.Selectuser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("程序运行端口", ini.Port)
	_ = r.Run(":" + strconv.Itoa(ini.Port))
}
