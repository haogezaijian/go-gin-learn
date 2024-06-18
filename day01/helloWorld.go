package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {

	//创建路由
	r := gin.Default()

	//绑定路由规则，执行函数
	//gin.Context,封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello World")
	})

	//监听端口，默认8080
	r.Run(":8000")
}