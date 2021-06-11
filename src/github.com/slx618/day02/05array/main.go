package main

import "fmt"

//数组
//必须指定存放的元素的类型和长度
//长度是数组类型的一部分 类型不同就不能用来比较
func main() {
	var arr1 [3]bool
	var arr2 [4]bool

	fmt.Printf("%T %T \n", arr1, arr2)

	// 数组的初始化

	//如果不初始化 默认元素都是零值
	fmt.Println(arr1, arr2)

	// 1. 初始化方式 1
	arr1 = [3]bool{true, true, true}
	// 2. 初始化方式 2

	//跟进初始值自动推断
	arr3 := [...]int{1, 2, 3, 5}
	fmt.Println(arr3)

	//初始化 3 根据索引初始化
	arr4 := [5]int{3: 1, 4: 2}
	fmt.Println(arr4)

	// 数组遍历
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	for k, _ := range arr4 {
		fmt.Println(arr4[k])
	}

	// 多维数组
	var arr5 [2][5]int
	arr5 = [2][5]int{
		[5]int{2, 2, 3, 5, 8},
		[5]int{333, 323, 32, 3, 2},
	}
	arr5[0][4] = 123

	fmt.Println(arr5)
	for _, v := range arr5 {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
