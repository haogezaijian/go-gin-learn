package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

/**
各种数据格式响应
 */

func main() {
	r := gin.Default()
	//json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"someJson","status":200})
	})
	//结构体
	r.GET("someStruct", func(c *gin.Context) {
		var msg struct{
			Name string
			Message string
			Number int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200,msg)
	})
	//XML
	r.GET("someXML", func(c *gin.Context) {
		c.XML(200,gin.H{"message":"abc"})
	})
	//YAML
	r.GET("someYAML", func(c *gin.Context) {
		c.YAML(200,gin.H{"name":"zhangsan"})
	})
	//protobuf
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1),int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(200,data)
	})

	r.Run(":8000")

}
