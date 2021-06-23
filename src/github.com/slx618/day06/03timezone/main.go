package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //本地时间
	fmt.Println(now)

	//要解析成东八区
	d, _ := time.Parse("2006-01-02 15:04:05", "2021-06-18 08:00:00")
	fmt.Println(d)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//按照指定时区解析
	dd, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-23 10:31:00", loc)
	fmt.Println(dd)

	fmt.Println(now.Sub(dd))
}
