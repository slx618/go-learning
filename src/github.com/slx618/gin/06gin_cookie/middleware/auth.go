package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Login struct {
	User      string   `json:"user"`
	RightList []string `json:"rightList"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		url := c.Request.URL.String()
		url = strings.Trim(url, "/")
		if url == "login" {
			c.Next()
			return
		}

		loginCookie, err := c.Cookie("login")
		if err != nil {
			log.Println("鉴权, 获取 cookie失败:", err)
		}

		var loginS Login
		log.Println(loginCookie)
		err = json.Unmarshal([]byte(loginCookie), &loginS)

		if err == nil {
			var pass bool
			log.Println(loginS.RightList)
			log.Println(loginS.User)
			for _, right := range loginS.RightList {
				if right == url {
					pass = true
				}
			}
			log.Println("pass", pass)
			if pass {
				c.Next()
			} else {
				c.Abort()
			}
		} else {
			log.Println("鉴权, 绑定 cookie 对象出错", err)
			c.Abort()
		}

	}
}
