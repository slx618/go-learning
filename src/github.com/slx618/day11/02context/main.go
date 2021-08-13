package main

import (
	"fmt"
	"sync"
	"time"
)

func f() {
	defer wg.Done()
breakTag:
	for {
		fmt.Println("demo")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-notify:
			break breakTag
		default:
		}
	}

}

var wg sync.WaitGroup

var notify chan bool = make(chan bool, 1)

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify <- true
	wg.Wait()
}
