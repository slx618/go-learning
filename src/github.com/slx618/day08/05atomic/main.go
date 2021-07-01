package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

func main() {

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&x, 1)
		}()
	}
	wg.Wait()
	fmt.Println(x)

	//比较成功后 替换成 12
	ok := atomic.CompareAndSwapInt64(&x, 1000000, 12)
	fmt.Println(x, ok)

}
