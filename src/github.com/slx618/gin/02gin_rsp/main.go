package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {

	r := gin.Default()

	// json 响应
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "success"})

	})

	r.GET("/struct", func(c *gin.Context) {
		var msg struct {
			Name string
			Msg  string
		}

		msg.Name = "xxx"
		msg.Msg = "success"

		c.JSON(200, msg)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"sss": 123})

	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"ss": "0ppp"})

	})

	r.GET("/protoBuf", func(c *gin.Context) {
		rps := []int64{1, 2}
		label := "label"

		//protoBuff
		data := &protoexample.Test{
			Label: &label,
			Reps:  rps,
		}

		c.ProtoBuf(200, data)

	})
	r.Run(":8090")
}
