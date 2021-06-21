package main

import "fmt"

/**
接口定义
type 接口名 interface {
	方法名(参数1, 参数2 ...) (返回值1, 返回值2)
	方法名(参数1, 参数2 ...) (返回值1, 返回值2)
	方法名(参数1, 参数2 ...) (返回值1, 返回值2)
	.
	.
	.
}
*/
type car interface {
	run()
}

type fll struct {
	brand string
}

func (f fll) run() {
	fmt.Println(f.brand, "开起来了")
}

type bsj struct {
	brand string
}

func (b bsj) run() {
	fmt.Println(b.brand, "开起来了")
}

func drive(c car) {
	c.run()
}

func main() {
	f1 := fll{"法拉利"}
	drive(f1)

	b1 := bsj{"保时捷"}
	drive(b1)

}
