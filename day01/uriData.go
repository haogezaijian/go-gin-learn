package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User1 struct {
	//binding:"required"修饰的字段，诺接收为空，则报错，是必须字段
	User 		string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password	string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/:user/:password", func(c *gin.Context) {
		u := &User1{}

		if err := c.ShouldBindUri(u); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		if u.User != "root" || u.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(200,gin.H{"status":"200"})
	})
	r.Run(":8000")
}
