package custom

import (
	"GinBlog/packages/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := context.Query("token")
		if token == "" {
			code = 400
		} else {
			cliams, err := jwt.ParseToken(token)
			if err != nil {
				code = 20001
			} else if time.Now().Unix() > cliams.ExpiresAt {
				code = 2002
			}
		}

		if code != 200 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"data": data,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
