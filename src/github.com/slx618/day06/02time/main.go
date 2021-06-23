package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//now.UTC()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Weekday())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time unix
	rst := time.Unix(1624415488, 0)
	fmt.Println(rst)
	fmt.Println(rst.Year())
	fmt.Println(rst.Second())
	fmt.Println(rst.Format("2006-01-02 15:04:05"))

	//时间间隔
	fmt.Println(time.Second)
	fmt.Println(time.Minute)

	//now + 1小时
	fmt.Println(now.Add(10 * time.Hour).String())

	//now.Before()
	//now.Sub()
	//now.Equal() 有时区

	fmt.Println("time sub", now.Sub(rst))
	//dd, _ := time.ParseDuration(now.Sub(rst).String())
	//fmt.Println("parse: ", dd)

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)
	//}

	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	fmt.Println(now.Format("2006年01月02日 15:04:05 PM"))
	fmt.Println(now.Format("2006年01月02日 15:04:05.000"))

	tp, _ := time.Parse("2006年01月02", "2021年06月23")

	fmt.Println(tp.Unix())

	//sleep
	//time.Sleep(time.Second)
	n := 1
	time.Sleep(time.Duration(n))
	time.Sleep(time.Second * 10)
}
