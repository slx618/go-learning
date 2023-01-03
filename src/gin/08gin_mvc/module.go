package main

type book struct {
	Id       int    `db:"id" form:"id"`
	BookName string `db:"book_name" form:"bookName"`
	Price    int    `db:"price" form:"price"`
}
