package main

import "fmt"

func main() {

	//中文字符 rune 英文字符 byte 属于类型别名
	s1 := "123abc我是小米"

	//for i := 0; i < len(s1); i++ {
	//	fmt.Println(s1[i])
	//	fmt.Printf("%c\n", s1[i])
	//}

	for _, v := range s1 {
		fmt.Printf("%c\n", v)
	}

	// "hello" => 'h' 'e' 'l' 'l' 'o'
	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //[白 萝 卜] 转换成 rune 切片
	//s3[0] = "红"
	s3[0] = '红'
	//转成字符串
	fmt.Println(string(s3))

	fmt.Printf("%T %T\n", s2, s3[0])
	fmt.Printf("%d\n", s3[0])
	fmt.Printf("%d\n", 'a')
	fmt.Printf("%d\n", 'A')

	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("%T\n", f)

	num := fmt.Sprintf("%d", s3[0])
	fmt.Printf(string(num))

}
