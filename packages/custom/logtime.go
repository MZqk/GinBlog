package custom

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logtime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context实例设置一个值
		//c.Set("geektutu", "1111")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(c.FullPath(), "请求耗时", latency)
	}
}
