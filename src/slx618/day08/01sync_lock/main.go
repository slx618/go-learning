package main

import (
	"fmt"
	"sync"
)

var x = 0

var (
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock() //操作先加🔐
		x += 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go add() // 反复赋值导致 x + 1 的值丢失 所以要加锁

	wg.Wait()
	fmt.Println(x)
}
