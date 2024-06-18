package main

import "github.com/gin-gonic/gin"

/**
routes group是为了管理一些相同的URL

httproter会将所有路由规则构造一颗前缀树
 */

func main() {
	//创建路由
	r := gin.Default()

	//路由组1，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}
	v2 :=r.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "haoge")
	c.String(200,name)
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "haoge")
	c.String(200,name)
}