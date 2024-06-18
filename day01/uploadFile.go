package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
multipart/form-data格式用于文件上传
gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
 */

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file,err := c.FormFile("file")
		if err != nil {
			c.String(500,"上传失败")
		}
		c.SaveUploadedFile(file,file.Filename)
		c.String(http.StatusOK,file.Filename)
	})
	r.Run()
}