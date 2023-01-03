package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
)

//下载第三方包 默认保存在 gopath 下面的 src
//go get -u github.com/go-sql-driver/mysql
func main() {

	dsn := "root:rV1pfnsuMxu3wxbu@tcp(127.0.0.1:3306)/xxl_job"
	db, err := sql.Open("mysql", dsn) // 不会校验用户名或者密码正确
	if err != nil {                   //这里检查的是 dsn 格式是否正确
		fmt.Println(db)
		return
	}
	defer db.Close()

	err = db.Ping() //尝试链接数据库
	if err != nil {
		fmt.Println(db)
		return
	}
}
