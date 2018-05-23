package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/zhihu-golang-web/models"
	"github.com/zhihu-golang-web/pkg/setting"
	"github.com/zhihu-golang-web/pkg/util"
	"net/http"
	"github.com/zhihu-golang-web/pkg/msg"
)

func GetProduct(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code :=msg.SUCCESS
	if id < 1{
		code = msg.INVALID_PARAMS
	}

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["product"] = models.GetProduct(id)
	c.JSON(http.StatusOK, gin.H{
		"code" :code,
		"data": data,
		"msg":msg.GetMsg(code),
	})
}

func GetProducts(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := msg.SUCCESS
	data["lists"] = models.GetProducts(util.GetPage(c), setting.App_.PAGE_SIZE, maps)
	data["total"] = models.GetProductTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}

//func DeleteArticle(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//
//	code := msg.INVALID_PARAMS
//
//
//		models.DeleteArticle(id)
//		code = e.SUCCESS
//	} else {
//		code = e.ERROR_NOT_EXIST_ARTICLE
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
