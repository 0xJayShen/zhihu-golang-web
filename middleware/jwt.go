package middleware

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = msg.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = msg.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != msg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
