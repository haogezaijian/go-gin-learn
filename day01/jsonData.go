package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
json数据解析和绑定
客户端传参，后端接收并解析到结构体
 */

type Login struct {
	//binding:"required"修饰的字段，诺接收为空，则报错，是必须字段
	User 		string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password	string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Password != "admin"{
			c.JSON(http.StatusBadRequest,gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	r.Run(":8000")
}
