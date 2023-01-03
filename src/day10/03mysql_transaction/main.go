package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //连接池

type user struct {
	id       uint64
	username string
	password string
}

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
	//设置数据库连接最大闲置链接数 会把多余连接关闭了
	db.SetMaxIdleConns(3)
	return err
}

func ts() error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	sql := `UPDATE user SET username=? WHERE id = ?`
	rst, err := db.Exec(sql, "武当山", 2)
	if err != nil {
		tx.Rollback()
		return err

	}
	last1, _ := rst.RowsAffected()
	if last1 == 0 {
		tx.Rollback()
		return errors.New("更新失败1")
	}

	rst1, err := db.Exec(sql, "武当四", 3)
	if err != nil {
		tx.Rollback()
	}

	last2, _ := rst1.RowsAffected()
	if last2 == 0 {
		tx.Rollback()
		return errors.New("更新失败2")
	}
	tx.Commit()

	fmt.Println(last1, last2)

	return err
}

//mysql事务
func main() {
	_ = initDb()
	err := ts()
	if err != nil {
		fmt.Println(err)
	}
}
