package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"zhihu-golang-web/models"
	"zhihu-golang-web/pkg/setting"
	"zhihu-golang-web/pkg/util"
	"net/http"
	"zhihu-golang-web/pkg/msg"
)

func GetProducts(c *gin.Context){
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
	code := e.SUCCESS
	data["lists"] = models.GetProducts(util.GetPage(c), setting.App_.PAGE_SIZE, maps)
	data["total"] = models.GetProductTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg.GetMsg(code),
		"data" : data,
	})
}
