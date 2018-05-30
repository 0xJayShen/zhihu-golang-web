package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/asdfsx/zhihu-golang-web/pkg/msg"
	"github.com/asdfsx/zhihu-golang-web/models"
	"github.com/asdfsx/zhihu-golang-web/pkg/util"
)

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := msg.INVALID_PARAMS

	isExist := models.CheckAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = msg.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = msg.SUCCESS
		}

	} else {
		code = msg.ERROR_AUTH
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}
