package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
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

func queryOne(id int) {
	var u user
	sql := `SELECT id,username,password FROM user WHERE id=?`
	//db.QueryRow(sql, 2)

	//scan 会调用释放 连接
	err := db.QueryRow(sql, id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", u)
}

func queryAll() {
	sql := `SELECT id,username,password FROM user`
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var u user
	var m map[int]user
	m = make(map[int]user, 10)
	for i := 0; rows.Next(); i++ {
		rows.Scan(&u.id, &u.username, &u.password)
		m[i] = u
	}
	fmt.Println(m)
}

func insert() {
	sql := `INSERT INTO user(username, password) VALUES(?, ?)`
	rst, err := db.Exec(sql, time.Now().Format("2006-01-02 15:04:05"), time.Now().UnixNano())
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ := rst.LastInsertId()
	fmt.Println("last insert id:", id)

}

//预处理插入
func prepareInsert() {
	sql := `INSERT INTO user(username, password) VALUES(?, ?)`

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rst, err := stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), time.Now().UnixNano())
	if err != nil {
		fmt.Println(err)
		return
	}

	id, _ := rst.LastInsertId()
	fmt.Println("last insert id:", id)

}

func updateRow(name, pass string, id int) {
	sql := `UPDATE user set username=?, password=? WHERE id=?`
	rst, err := db.Exec(sql, name, pass, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	eff, _ := rst.RowsAffected()
	fmt.Println("effort row:", eff)

}

func delRow(id int) {
	sql := `DELETE FROM user WHERE id = ?`
	rst, err := db.Exec(sql, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	effId, _ := rst.RowsAffected()
	fmt.Println("effort row:", effId)
}

func main() {

	err := initDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	//queryOne(3)
	//queryAll()

	//insert()
	//updateRow("武当王也", "232323232", 3)
	//delRow(11)
	prepareInsert()
	prepareInsert()
	prepareInsert()
	prepareInsert()
}
