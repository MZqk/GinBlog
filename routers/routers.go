package routers

import (
	"GinBlog/models"
	"GinBlog/packages/ini"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Routers() {
	r := gin.Default()
	r.GET("/create", func(context *gin.Context) {
		models.Createorm()
		context.JSON(http.StatusOK, gin.H{
			"datas":   "创建数据成功",
			"webroot": "create",
		})
	})
	r.GET("/select", func(context *gin.Context) {
		datas := models.Selectorm()
		context.IndentedJSON(http.StatusOK, gin.H{
			"datas":   datas,
			"webroot": "select",
		})
	})

	fmt.Println("程序运行端口", ini.Port)
	_ = r.Run(":" + strconv.Itoa(ini.Port))
}
