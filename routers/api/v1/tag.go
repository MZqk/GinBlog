package v1

import (
	"GinBlog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary 新疆tag
// @Description 上传tag的name
// @Produce  json
// @Param  ID query string true "id"
// @Success 200 {string} json {"data":{},"webroot":""}
// @Router /api/v1/tag/ [post]
func Createtag(ctx *gin.Context) {
	tag := ctx.Query("tag")
	models.Createtag(tag)
	ctx.JSON(http.StatusOK, gin.H{
		"datas":   "创建数据成功",
		"webroot": "tag",
	})
}

// @Summary 查询tag
// @Description 根据tag的id查询tag的name
// @Produce  json
// @Param  ID query string true "id"
// @Success 200 {string} json {"data":{},"webroot":""}
// @Router /api/v1/tag/ [get]
func Selecttag(ctx *gin.Context) {
	datas := models.Selecttag()
	ctx.YAML(http.StatusOK, gin.H{
		"datas":   datas,
		"webroot": "tag",
	})
}

func Gettag(ctx *gin.Context) {
	tagid := ctx.Param("id")
	id, _ := strconv.Atoi(tagid)
	fmt.Println(id)
	datas := models.Gettag(id)
	time.Sleep(time.Second * 2)
	ctx.JSON(http.StatusOK, gin.H{
		"datas":   datas,
		"webroot": "tag",
	})

}
