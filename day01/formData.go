package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
表单数据解析和绑定
 */
type User struct {
	//binding:"required"修饰的字段，诺接收为空，则报错，是必须字段
	User 		string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password	string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/loginForm", func(c *gin.Context) {
		var form User
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	r.Run(":8000")
}
