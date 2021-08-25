package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	//异步
	r.GET("/long_async", func(c *gin.Context) {
		//异步的情况下要用副本
		cc := c.Copy()
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步:", cc.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步:", c.Request.URL.Path)
	})

	r.Run(":8090")
}
