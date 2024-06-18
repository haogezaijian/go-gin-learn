package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
中间件练习
 */

func myTime(c *gin.Context)  {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("程序用时：",since)
}

func main() {
	r := gin.Default()

	r.Use(myTime)

	shopingGroup := r.Group("/shopping")
	{
		shopingGroup.GET("/index",index)
		shopingGroup.GET("/home",home)
	}
	r.Run()
}

func index(c *gin.Context) {
	time.Sleep(5*time.Second)
}

func home(c *gin.Context) {
	time.Sleep(3*time.Second)
}
