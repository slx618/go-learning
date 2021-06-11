package main

import "fmt"

func main() {

	s1 := []string{"北京", "上海", "广州"}
	//s1[3] = "深圳" //索引越界
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "深圳")
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "重庆", "南京", "杭州")
	ss := s1
	// 底层数组追加但是放不下的时候会重新申请内存
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	//ss 拆开传入
	s1 = append(s1, ss...)
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

}
