package app

import (
	"github.com/gin-gonic/gin"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
)
type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  msg.GetMsg(errCode),
		"data": data,
	})

	return
}
