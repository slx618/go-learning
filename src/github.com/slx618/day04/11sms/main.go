package main

import (
	"fmt"
	"os"
)

/**
能够查询 新增 删除 学生
*/
var (
	allStu map[int64]*student
)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

func showAllStu() {
	if len(allStu) > 0 {
		for _, v := range allStu {
			fmt.Println("学号:", v.id, "姓名:", v.name)
		}
	} else {
		fmt.Println("没有学员")
	}
}

func addStu() {
	var (
		id   int64
		name string
	)
checkId:
	fmt.Print("请输入学号:")
	fmt.Scan(&id)
	for k := range allStu {
		if k == id {
			fmt.Println("学号已存在")
			goto checkId
		}
	}
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)
	s := newStudent(id, name)
	allStu[id] = &s
	fmt.Println("添加成功")

}

func delStu() {
	var id int64
	fmt.Print("请输入要删除的学号:")
	fmt.Scan(&id)
	delete(allStu, id)
}

func main() {
	allStu = make(map[int64]*student, 100)
	for {
		//1. 打印菜单
		fmt.Println("欢迎~~~~~~")
		fmt.Println(`
		1. 查看学生
		2. 新增学生
		3. 删除学生
		4. 退出`)
		fmt.Print("请输入:")
		var ch int
		fmt.Scanln(&ch)
		//2. 待用户选择

		fmt.Printf("选择了: %d\n", ch)
		//3. 执行对应功能
		switch ch {
		case 1:
			showAllStu()
		case 2:
			addStu()
		case 3:
			delStu()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
