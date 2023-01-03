package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		r := rand.Intn(10)
		fmt.Println(r)

	}

}
func f1(i int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()
	//wg.Add(10)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
