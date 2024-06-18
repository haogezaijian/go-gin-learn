package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
中间件
全局中间件，所有请求都经过此中间件
 */

/**
中间件（Middleware）在Web开发中扮演着非常重要的角色，它通常位于应用程序的请求处理管道中，用于处理进入应用程序的HTTP请求和离开应用程序的HTTP响应。
中间件提供了一种机制，可以在请求到达目标路由处理程序之前或之后执行特定的任务或操作。
 */


//定义之间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		//设置变量到Context的key中，可以通过Get()取
		c.Set("request","中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完成", status)
		t2 := time.Since(t)
		fmt.Println(t2)
	}
}
func main() {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare())
	{
		r.GET("/ce", func(c *gin.Context) {
			req,_ := c.Get("request")
			fmt.Println("request:",req)
			c.JSON(200,gin.H{"request":req})
		})
	}
	r.Run()
}
