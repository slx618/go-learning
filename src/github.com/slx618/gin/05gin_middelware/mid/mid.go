package mid

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("middle start")
		//设置变量到 context
		c.Set("mid", "中间件")
		//执行函数
		c.Next() //执行了 next 后请求实际上已经分发到具体的路由上去了
		//响应
		status := c.Writer.Status()
		t2 := time.Since(t)
		log.Println("middle end", status)
		log.Println("耗时", t2)

	}
}

func M2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("middle 🐴 start")
		//设置变量到 context
		c.Set("mid", "中间件")
		//执行函数
		c.Next() //执行了 next 后请求实际上已经分发到具体的路由上去了
		//响应
		status := c.Writer.Status()
		t2 := time.Since(t)
		log.Println("middle 🐴 end", status)
		log.Println("耗时", t2)

	}
}
