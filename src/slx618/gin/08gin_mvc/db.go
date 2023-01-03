package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var Db *sqlx.DB

func initDb() (err error) {
	dsn := "root:rV1pfnsuMxu3wxbu@tcp(127.0.0.1:3306)/xxl_job"
	Db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Println("mysql init err:", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		log.Println("mysql init ping err:", err)
		return
	}

	Db.SetMaxIdleConns(5)
	Db.SetMaxOpenConns(50)
	Db.SetConnMaxLifetime(time.Second * 3)
	return err
}

func getBookList() (bookList []book) {
	s := "SELECT id,book_name,price FROM book"
	rows, err := Db.Query(s)
	if err != nil {
		log.Println("query data err", err)
		return
	}

	defer rows.Close()
	var b book
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&b.Id, &b.BookName, &b.Price)
		bookList = append(bookList, b)
	}
	return
}

func addBook(bookInfo book) bool {
	s := "INSERT INTO book(book_name, price, create_time) VALUES(?, ?, ?)"
	_, err := Db.Exec(s, bookInfo.BookName, bookInfo.Price, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("insert data err", err)
		return false
	}
	return true
}

func delBook(id string) bool {
	s := "DELETE FROM book WHERE id = ?"
	_, err := Db.Exec(s, id)
	if err != nil {
		log.Println("del data err", err)
		return false
	}
	return true
}
