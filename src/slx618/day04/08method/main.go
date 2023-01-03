package main

import "fmt"

//标识符: 变量名, 函数名 类型名 方法名
//go语言中如果标识符是大写的 表示对外部包可见

// Dog 是一个狗的结构体
//type Dog struct {
//	name string
//}

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string) person {
	return person{
		name: name,
	}
}

//使用值接受: 传拷贝进去
func (p person) guonian() {
	p.age++
}

//使用指针接收
//要改变值的时候
//大对象的时候
//保证一致性 一 个方法使用了指针接收者 其他方法也尽量使用指针接收者
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {

}

//方法是作用于特定类型的函数
//接受者表示是调用该方法的具体类型变量, 多用类型字母小写
func (d dog) wang() {
	fmt.Printf("%s: wangwangwang\n", d.name)
}

func (p person) wang() {
	fmt.Printf("%s: wangwangwang\n", p.name)
}

func main() {
	d1 := newDog("阿黄")
	d1.wang()

	p1 := newPerson("xxx")
	p1.wang()
	p1.guonian()
	fmt.Println(p1)
	p1.zhenguonian()
	fmt.Println(p1)

}
