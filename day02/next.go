package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
Next方法
 */
/**
在 Gin Web 框架中，c.Next() 是一个非常重要的方法，用于在中间件（Middleware）中调用。它允许你执行当前中间件之后的处理逻辑，这通常包括后续的中间件（如果有的话）和最终的路由处理函数（Handler）。

当你编写一个中间件时，你可能会希望执行一些前置逻辑（比如日志记录、身份验证、授权等），然后调用 c.Next() 来继续处理请求。一旦 c.Next() 被调用，Gin 将会执行在请求处理链中的下一个中间件或路由处理函数。

在 c.Next() 被调用之后，你还可以执行一些后置逻辑，这些逻辑将在请求处理链中的所有中间件和路由处理函数都执行完毕后执行。这通常用于记录请求处理的总时间、清理资源、发送额外的响应头等。
 */
// 定义中间
func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare2())
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run()
}
