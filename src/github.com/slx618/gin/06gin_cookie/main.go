package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/slx618/gin/06gin_cookie/middleware"
	"log"
)

func main() {

	r := gin.Default()

	r.Use(middleware.Auth())
	{
		r.GET("/login", func(c *gin.Context) {
			var login middleware.Login
			login.User = "admin"
			login.RightList = append(login.RightList, "home")
			log.Println(login)
			cookie, err := json.Marshal(login)
			if err != nil {
				log.Println("json encode err", err)
			}

			c.SetCookie("login", string(cookie), 300, "/", "localhost", true, true)
			c.String(200, "success")
		})

		r.GET("/home", func(c *gin.Context) {
			c.String(200, "this is home")
		})

	}

	r.Run(":8091")

}
