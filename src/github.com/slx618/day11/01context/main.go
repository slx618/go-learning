package main

import (
	"fmt"
	"sync"
	"time"
)

func f() {
	defer wg.Done()
	for {
		fmt.Println("demo")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}

}

var wg sync.WaitGroup

var notify bool

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
