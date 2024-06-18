package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
服务端发送cookie给客户端，客户端请求时携带cookie
 */

func main() {
	r := gin.Default()
	r.GET("cookie", func(c *gin.Context) {
		cookie,err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Println("cookie 的值：",cookie)
	})
	r.Run()
}
