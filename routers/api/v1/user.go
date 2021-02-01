package v1

import (
	"GinBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Createuser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	models.Createuser(username, password)
	ctx.JSON(http.StatusOK, gin.H{
		"datas":   "用户创建成功",
		"webroot": "user",
	})
}
func Selectuser(ctx *gin.Context) {
	datas := models.Selectuser()
	ctx.JSON(http.StatusOK, gin.H{
		"datas":   datas,
		"webroot": "user",
	})
}
