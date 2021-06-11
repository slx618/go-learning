package main

import (
	"fmt"
	"strings"
)

//字符串
//s := "字符串"

//字符
// c1 := 'h'
// c1 := '1'
// c1 := '字'

//字节 1 字节 = 1 byte = 8 bit 八个二进制
//1 个字符 'A' = 1个字节
// 1个 utf8 编码的汉字 一般三个字节
func main() {
	fmt.Println("\\\"")

	s2 := `第一行
	第二行
	`
	fmt.Println(s2)

	s3 := `\`
	fmt.Println(s3)
	// len() 长度
	fmt.Println(len(s2))

	// 字符串拼接
	name := "小米"
	make := "煮水"
	ss := name + make
	fmt.Println(ss)
	ss1 := fmt.Sprintf(ss)
	fmt.Println(ss1)

	s4 := "/Users/slx618/github/go-movies/src/github.com/slx618/day01/08string"

	rst := strings.Split(s4, "/")
	fmt.Println(rst)

	fmt.Println(strings.Contains(s4, "slx618"))
	fmt.Println(strings.Contains(s4, "000"))
	fmt.Println(strings.HasPrefix(s4, "U"))
	fmt.Println(strings.HasSuffix(s4, "g"))
	fmt.Println(strings.Index(s4, "g"))     //第一次出现的
	fmt.Println(strings.LastIndex(s4, "g")) //最后一次出现的地方

	fmt.Println(strings.Join(rst, "+"))
	fmt.Println(strings.Replace(s4, "/", "+", 2))

}
