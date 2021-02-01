package routers

import (
	"GinBlog/models"
	"GinBlog/packages/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string
	Password string
}

func GetAuth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//a := auth{Username: username,Password: password}

	data := make(map[string]interface{})
	var code int
	if models.CheckAuth(username, password) {
		token, err := jwt.GenerateToken(username, password)
		if err != nil {
			code = 400
		} else {
			data["token"] = token
			code = 200
		}
	} else {
		code = 401
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": code,
	})
}
