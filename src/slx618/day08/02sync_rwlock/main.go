package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var lock sync.Mutex
var rwLock sync.RWMutex
var wg sync.WaitGroup

func main() {
	begin := time.Now()
	for i := 0; i < 100; i++ {
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		go read()

	}

	wg.Wait()
	fmt.Println(time.Now().Sub(begin))
}
func read() {
	wg.Add(1)
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 20)
	rwLock.RUnlock()
}

func write() {
	wg.Add(1)
	defer wg.Done()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 30)
	rwLock.Unlock()

}
