package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qq976739120/zhihu-golang-web/pkg/setting"
	"github.com/qq976739120/zhihu-golang-web/routers/v1"
	//"zhihu-golang-web/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode_.RUN_MODE)
	r.GET("/auth", GetAuth)
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
	{
		//获取标签列表
		apiv1.GET("/products", v1.GetProducts)
		apiv1.GET("/product/:id", v1.GetProduct)

		//新建标签
	}
	return r
}
