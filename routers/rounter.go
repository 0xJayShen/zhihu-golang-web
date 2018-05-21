package routers

import (
	"github.com/gin-gonic/gin"
	"zhihu-golang-web/pkg/setting"
	"zhihu-golang-web/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode_.RUN_MODE)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/products", v1.GetProducts)
		//新建标签

	}
	return r
}
