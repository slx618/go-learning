package main

import (
	"flag"
	"fmt"
)

func main() {
	//name := flag.String("name", "luo", "请输入姓名")
	//age := flag.Int("age", 18, "请输入年龄")
	//dur := flag.Duration("dur", time.Hour, "结婚年龄")
	////使用
	//flag.Parse()
	//fmt.Println(*name)
	//fmt.Println(*age)
	//fmt.Println(*dur)

	var name string
	flag.StringVar(&name, "name", "luo", "请输入姓名")
	flag.Parse() // 这行一定要有啊~~~~~~~!!!!
	fmt.Println(name)

	fmt.Println(flag.Args()) //命令行参数以[]string的形式返回
	fmt.Println(flag.NArg()) // 返回命令行参数的总数
	fmt.Println(flag.NFlag())
}
