package main

import "fmt"

func main() {

	fmt.Print("不会自动换行\n")
	fmt.Println("会自动换行\n")
	fmt.Println("行行行行行行行")

	fmt.Printf("%v\n", "ssssss")
	fmt.Printf("%#v\n", "ssssss")
	//%#v 会打印 值的go语法表示
	//%+v 结构体输出
	//%v 值原样输出
	//%T 查看类型

	//%d 10 进制数
	//%f %F 浮点数
	// 宽度标识符 %2f %.2f
	//%e $E
	//%g %G 自动使用 %f %e

	//%b 2进制
	//%o 8进制
	//%x 16进制
	//%X 16进制 大写的字母

	//%c 字符Unicode码值
	//%C 字符Unicode 格式

	//%s 字符串 %2s 占位符不足 补空格 %-2s
	//%p 指针
	//%q ascii 码转成值 | 把字符串包起来
	//%% 百分号
	//%t 布尔值

	var m1 = make(map[string]int, 1)
	m1["ssss"] = 1
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)

	fmt.Printf("%T\n", m1)
	fmt.Printf("%q\n", 1)

	fmt.Printf("%9f\n", 1.2)
	fmt.Printf("%4s\n", "刷刷刷")
	fmt.Printf("%-2s\n", "刷")
	//.点后面截取
	fmt.Printf("%1.2s\n", "刷刷刷")

	//获取用户输入 fmt.Scan
	//var s string
	//fmt.Scan(&s)
	//fmt.Println(s)
	//获取用户输入 fmt.Scanf
	var (
		name  string
		age   int
		class string
	)
	//	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//	fmt.Println(name, age, class)

	//获取用户输入 fmt.Scanln
	fmt.Scanln(&name, &age, &class)

	fmt.Println(name, age, class)

}
