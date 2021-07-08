package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //连接池

func initDb() (err error) {
	dsn := "root:rV1pfnsuMxu3wxbu@tcp(127.0.0.1:3306)/xxl_job"
	db, err = sql.Open("mysql", dsn) // 注意不要用 := 否则 db 成了局部变量
	if err != nil {                  //这里检查的是 dsn 格式是否正确
		return err
	}

	err = db.Ping() //尝试链接数据库
	if err != nil {
		return err
	}

	//设置数据连接池 最大链接数
	db.SetMaxOpenConns(10)
	return err
}

type user struct {
	id       uint64
	username string
	password string
}

func main() {

	err := initDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	var u user
	sql := `SELECT id,username,password FROM user WHERE id=?`
	//db.QueryRow(sql, 2)
	err = db.QueryRow(sql, 2).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)

}
