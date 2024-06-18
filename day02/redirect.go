package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"/redirect")
	})
	r.GET("/redirect", func(c *gin.Context) {
		c.JSON(200, map[string]string{"name":"haoge"})
	})
	r.Run()

}
