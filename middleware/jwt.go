package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
	"time"
	"net/http"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
	"fmt"
	"github.com/qq976739120/zhihu-golang-web/cache"
	"github.com/garyburd/redigo/redis"
)

//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code int
//		var data interface{}
//
//		code = msg.SUCCESS
//		token := c.Query("token")
//
//		if token == "" {
//			code = msg.INVALID_PARAMS
//		} else {
//			claims, err := util.ParseToken(token)
//			if err != nil {
//				code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
//			} else if time.Now().Unix() > claims.ExpiresAt {
//				code = msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
//			}
//		}
//
//		if code != msg.SUCCESS {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"code": code,
//				"msg":  msg.GetMsg(code),
//				"data": data,
//			})
//
//			c.Abort()
//			return
//		}
//
//		c.Next()
//	}
//}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = msg.SUCCESS
		cookie, err := c.Request.Cookie("token")

		if err != nil {
			fmt.Println(err)
		}
		if cookie != nil {
			token := cookie.Value

			if token == "" {
				code = msg.INVALID_PARAMS
			} else {
				conn := cache.RedisPool.Get()
				defer conn.Close()
				exit_token_redis, _ := redis.Bool(conn.Do("EXISTS", token))
				if exit_token_redis {
					token_redis, _ := redis.String(conn.Do("GET", token))
					if token == token_redis {
						claims, err := util.ParseToken(token)
						if err != nil {
							code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
						} else if time.Now().Unix() > claims.ExpiresAt {
							code = msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
						}
					}
				}
			}
		} else {
			code = msg.INVALID_PARAMS
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
