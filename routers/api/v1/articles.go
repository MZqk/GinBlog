package v1

import (
	"GinBlog/models"
	"github.com/gin-gonic/gin"
)

func CreateArticle(ctx *gin.Context) {
	context := ctx.PostForm("context")
	models.CreateArticle(context)
}
