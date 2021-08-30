package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	err := initDb()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.LoadHTMLGlob("tpl/*")
	r.GET("/bookList", func(c *gin.Context) {
		data := getBookList()
		c.HTML(200, "bookList.html", gin.H{
			"code": 200,
			"data": data,
		})
	})

	r.GET("/addBook", func(c *gin.Context) {
		c.HTML(200, "addForm.html", nil)

	})

	r.POST("/addBook", func(c *gin.Context) {
		var book book
		if err := c.ShouldBind(&book); err == nil {
			rst := addBook(book)
			if rst {
				c.Redirect(http.StatusMovedPermanently, "/bookList")
			}
		} else {
			log.Println("参数错误", book)
		}
	})

	r.GET("/delBook", func(c *gin.Context) {

		id := c.DefaultQuery("id", "")
		log.Println(id)
		if id == "" {
			log.Println("参数错误")
		}

		rst := delBook(id)
		if rst {
			c.Redirect(302, "/bookList")
		}
	})

	r.Run(":8091")
}
