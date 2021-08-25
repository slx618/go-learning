package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type login struct {
	Name string `form:"user" json:"user" uri:"user" binding:"required"`
	Pwd  string `form:"pwd" json:"pwd" uri:"pwd" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()
	//绑定路由规则
	r.GET("/user/:name/*action", func(c *gin.Context) {
		controller := c.Param("name")
		action := c.Param("action")

		query := c.DefaultQuery("name", "jjj")
		c.String(200, "hello world: %s %s default: %s", controller, action, query)
	})

	r.POST("/post", func(c *gin.Context) {
		n := c.DefaultPostForm("name", "xx")
		pwd := c.PostForm("pwd")
		h := c.PostFormArray("like")
		if pwd == "" {
			c.String(401, "")
		}

		c.String(200, "%v + %v + %v", n, pwd, h)

	})

	//大小 默认为 32M
	r.MaxMultipartMemory = 8 << 20 //设置为 8M
	r.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Println(err)
		}
		c.String(200, "filename:%s, filesize:%s", file.Filename, file.Size)
	})

	r.POST("/mUpload", func(c *gin.Context) {
		f, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "err: %s", err)
		}

		files := f.File["file"]

		for _, file := range files {
			err = c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				fmt.Println(err)
			}
			c.String(200, "filename:%s, filesize:%d\r\n", file.Filename, file.Size)

		}

	})

	v1 := r.Group("/v1")
	{
		v1.GET("g", func(c *gin.Context) {
			c.String(200, "success")
		})
	}

	v2 := r.Group("v2")
	{
		v2.POST("g", func(c *gin.Context) {
			c.String(200, "post success")
		})
	}

	r.POST("json", func(c *gin.Context) {
		var login login

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error(), "status": false})
			return
		}

		c.JSON(200, login)

	})

	r.POST("form", func(c *gin.Context) {
		var login login

		//会处理
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error(), "status": false})
			return
		}

		c.JSON(200, login)

	})

	r.POST("/uri/:user/:pwd", func(c *gin.Context) {
		var uri login
		//会处理
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error(), "status": false})
			return
		}

		c.JSON(200, uri)

	})
	//监听端口
	r.Run(":8090")

}
