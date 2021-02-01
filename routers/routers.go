package routers

import (
	"GinBlog/packages/custom"
	"GinBlog/packages/ini"
	v1 "GinBlog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io"
	"os"
	"strconv"
)

func Routers() {
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// @title 测试
	// @version 0.0.1
	// @description  测试
	// @BasePath /api/v1/
	apiv1 := r.Group("/api/v1")
	apiv1.Use(custom.Logtime())
	apiv1.Use(custom.JWT())
	{
		apiv1.POST("/tag/", v1.Createtag)
		apiv1.GET("/tag/:id", v1.Gettag)
		apiv1.GET("/tags", v1.Selecttag)
		apiv1.POST("/user", v1.Createuser)
		apiv1.GET("/users", v1.Selectuser)
		apiv1.POST("/article", v1.CreateArticle)
		//apiv1.POST("/article",)
	}
	r.POST("/auth", GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//fmt.Println("程序运行端口", ini.Port)
	_ = r.Run(":" + strconv.Itoa(ini.Port))

	//使用GO 自带库het/http的新特性，减少频繁重启调试
	//srv := &http.Server{
	//	Addr:    ":" + strconv.Itoa(ini.Port),
	//	Handler: r,
	//}
	//go func() {
	//	// 服务连接
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//log.Println("Server exiting")
}
