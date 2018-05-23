package routers

import (
	"github.com/qq976739120/zhihu-golang-web/pkg/logging"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
	"github.com/qq976739120/zhihu-golang-web/models"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"

)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := msg.INVALID_PARAMS
	if ok {
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
	} else {
		for _, err := range valid.Errors {
			logging.Info(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}
