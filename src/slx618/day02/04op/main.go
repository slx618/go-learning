package main

import "fmt"

//运算符

func main() {
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ //单独的语句 不能放在 = 号右边
	b--
	fmt.Println(a)
	fmt.Println(b)

	//关系运算符
	fmt.Println(a == b) //相同类型的变量才能比较
	fmt.Println(a != b) //相同类型的变量才能比较
	fmt.Println(a >= b) //相同类型的变量才能比较
	fmt.Println(a <= b) //相同类型的变量才能比较
	fmt.Println(a > b)  //相同类型的变量才能比较
	fmt.Println(a < b)  //相同类型的变量才能比较

	//逻辑运算符
	var age = 10
	if age > 19 && age < 60 {
		fmt.Println("成了")
	}

	if age < 19 || age > 60 {
		fmt.Println("不行")
	}

	var isB bool = true
	fmt.Println(isB)
	fmt.Println(!isB)

	// 位运算符 针对二进制数
	//5 的二进制 101
	//2 的二进制 10

	//& 按位与 两位 都为 1 才是 1
	fmt.Println(5 & 2)

	// | 按位或 两位有一个为 1 就为 1
	fmt.Println(5 | 2)

	// ^ 异或 两位不一样 就为 1
	fmt.Println(5 ^ 2)

	// << 将二进制为左移指定位数
	fmt.Println(5 << 1)  // 101 => 1010 => 10
	fmt.Println(1 << 10) // 10000000000
	// >> 将二进制为右移指定位数
	fmt.Println(5 >> 1) // 101 => 10 => 2
	fmt.Println(5 >> 2) // 101 => 1 => 1
	fmt.Println(5 >> 3) // 0

	var m = int8(1)
	fmt.Println(m << 10) // 0  10000000000 取 8 位 00000000 => 0

	// 赋值运算
	var x = 10
	x += 1
	x -= 1
	x *= 1
	x /= 1
	x %= 1

	x <<= 2
	x >>= 2
	x &= 2
	x |= 2
	x ^= 2
}
