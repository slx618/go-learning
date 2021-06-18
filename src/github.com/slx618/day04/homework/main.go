package main

import (
	"fmt"
	"os"
)

type sms struct {
	stuList map[uint64]studemt
}

type studemt struct {
	stuNo   uint64
	stuName string
}

func newStu(stuNo uint64, stuName string) studemt {
	return studemt{
		stuNo:   stuNo,
		stuName: stuName,
	}

}

func (s sms) listStu() {
	if len(s.stuList) > 0 {
		for _, v := range s.stuList {
			fmt.Println("学号:", v.stuNo, "姓名:", v.stuName)
		}
	} else {
		fmt.Println("学员为空")
	}
}

func (s sms) addStu() {
	var (
		stuNo   uint64
		stuName string
	)
checkPoint:
	fmt.Print("请输入学号:")
	fmt.Scan(&stuNo)
	for k := range s.stuList {
		if k == stuNo {
			fmt.Println("学号冲突, 请重新输入")
			goto checkPoint
		}
	}

	fmt.Print("请输入姓名:")
	fmt.Scan(&stuName)
	fmt.Print("输入的学生姓名为:", stuName)

	stu := newStu(stuNo, stuName)
	s.stuList[stuNo] = stu
}

func (s sms) delStu() {
	var stuNo uint64

	fmt.Print("请输入要删除的学号")
	fmt.Scan(&stuNo)

	delete(s.stuList, stuNo)
}

func main() {
	s := sms{
		stuList: make(map[uint64]studemt, 9999),
	}
	for {

		fmt.Print(`
欢迎~~~~~~~~~~~
	请选择功能:
		1. 查看所有学生
		2. 添加学生 
		3. 删除学生
		4. 退出
`)
		fmt.Print("等待输入: ")
		var con int8
		_, err := fmt.Scan(&con)
		if err != nil {
			fmt.Println("输入有误:", err)
		}
		fmt.Println("你选择的是:", con)
		switch con {
		case 1:
			s.listStu()
		case 2:
			s.addStu()
		case 3:
			s.delStu()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("功能不存在")
		}
	}
}
