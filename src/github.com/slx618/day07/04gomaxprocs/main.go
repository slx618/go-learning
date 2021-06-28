package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Printf("a:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Printf("b:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(2) //不设置 默认逻辑核心数
	fmt.Println(runtime.NumCPU())
	go a()
	wg.Add(1)
	go b()
	wg.Add(1)
	wg.Wait()
}
