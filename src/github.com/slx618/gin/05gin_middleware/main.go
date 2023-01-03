package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slx618/gin/05gin_middleware/mid"
	"log"
)

//中间件
//必须是 gin.HandleFunc返回类型才能作为中间件
func main() {

	r := gin.Default()
	//注册中间件
	r.Use(mid.MiddleWare())
	{
		r.GET("/", func(c *gin.Context) {
			key, _ := c.Get("mid")
			log.Printf("响应请求%v\r\n", key)
			c.String(200, "success")
		})

		r.GET("/v", mid.M2(), func(c *gin.Context) {
			key, _ := c.Get("mid")
			log.Printf("响应请求%v\r\n", key)
			c.String(200, "success")
		})

	}

	r.Run(":8091")
}
