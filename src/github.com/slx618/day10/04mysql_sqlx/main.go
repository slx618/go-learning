package main

// go get -u github.com/jmoiron/sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

func main() {
	_ = initDb()

	var u user
	sql := `SELECT id,username FROM user WHERE id=?`
	err := db.Get(&u, sql, 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", u)

	sql = `SELECT id,username FROM user`
	var uList = make([]user, 0, 10) //长度 容量
	err = db.Select(&uList, sql)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", uList)

}
