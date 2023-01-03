package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//sql 注入

var db *sqlx.DB //连接池

type user struct {
	Id       uint64
	Username string
	Password string
}

func initDb() (err error) {
	dsn := "root:rV1pfnsuMxu3wxbu@tcp(127.0.0.1:3306)/xxl_job"
	db, err = sqlx.Connect("mysql", dsn) // 注意不要用 := 否则 db 成了局部变量
	if err != nil {                      //这里检查的是 dsn 格式是否正确
		return err
	}

	//设置数据连接池 最大链接数
	db.SetMaxOpenConns(10)
	//设置数据库连接最大闲置链接数 会把多余连接关闭了
	db.SetMaxIdleConns(3)
	return err
}

func sqlInject(name string) {
	sql := fmt.Sprintf(`SELECT * FROM user WHERE username="%s"`, name)

	fmt.Printf("sql: %s\n", sql)

	u := make([]user, 0, 10)
	err := db.Select(&u, sql)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ssss")
		return
	}

	fmt.Printf("%#v", u)
}

func main() {
	_ = initDb()

	//注入示例
	sqlInject("xxx\" OR 1=1 limit 1 #")
	sqlInject("xxx\" AND  (SELECT count(*) FROM user WHERE id < 10) #")
	sqlInject("xxx\" UNION  (SELECT * FROM user) #")
}
