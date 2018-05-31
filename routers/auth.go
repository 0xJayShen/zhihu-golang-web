package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
	"github.com/qq976739120/zhihu-golang-web/models"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
	"github.com/qq976739120/zhihu-golang-web/cache"
)

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	data := make(map[string]interface{})
	code := msg.INVALID_PARAMS

	user_id := models.CheckAuth(username, password)
	if user_id > 0  {
		conn := cache.RedisPool.Get()
		defer conn.Close()

		token, err := util.GenerateToken(username, password)
		conn.Do("SET",token,user_id)
		if err != nil {
			code = msg.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = msg.SUCCESS
		}

		cookie:=&http.Cookie{
			Name:   "token",
			Value:    token,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   msg.COOKIE_MAX_MAX_AGE,
		}

		http.SetCookie(c.Writer,cookie)

	} else {
		code = msg.ERROR_AUTH
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}
