package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
	"github.com/qq976739120/zhihu-golang-web/models"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
	"github.com/qq976739120/zhihu-golang-web/cache"
	"github.com/qq976739120/zhihu-golang-web/pkg/app"
	"google.golang.org/grpc/credentials/alts/core/conn"
)

func GetAuth(c *gin.Context) {

	appG := app.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")
	data := make(map[string]interface{})
	code := msg.INVALID_PARAMS

	user_id := models.CheckAuth(username, password)

	if user_id < 1  {
		appG.Response(http.StatusOK,msg.ERROR_AUTH_CHECK_TOKEN_FAIL,msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
		return
	}
	token, err := util.GenerateToken(username, password)
	if err!=nil {
		appG.Response(http.StatusOK,msg.ERROR_AUTH_TOKEN,msg.GetMsg(msg.ERROR_AUTH_TOKEN))

	}
	reply,err:=cache.Set(token,user_id,7*24*60*60)
	if reply!=true && err!= nil{
		appG.Response(http.StatusOK,msg.CACHE_SET_FAIL,msg.GetMsg(msg.CACHE_SET_FAIL))

	}





		token, err := util.GenerateToken(username, password)
		conn.Do("SET",token,user_id,"EX","2592000")
		if err != nil {
			appG
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
		//"data": data,
	})
}
