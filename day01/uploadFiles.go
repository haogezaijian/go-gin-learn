package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
上传多个文件
 */

func main() {

	r := gin.Default()
	//限制8M，默认为32M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form,err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest,fmt.Sprintf("get err %s",err.Error()))
		}
		files := form.File["files"]
		//遍历所以文件
		for _,file := range files {
			if err := c.SaveUploadedFile(file,file.Filename); err != nil {
				c.String(http.StatusBadRequest,fmt.Sprintf("upload err %s",err.Error()))
				return
			}
		}
		c.String(200,fmt.Sprintf("upload ok %d files",len(files)))
	})
	r.Run(":8000")

}
