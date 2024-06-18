package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
可以通过Context的Param方法来获取API参数
*/

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(action)
		action = strings.Trim(action,"/")
		c.String(http.StatusOK,name+"is"+action)
	})
	r.Run(":8000")
}
