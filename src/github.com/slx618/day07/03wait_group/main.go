package main

import (
	"fmt"
	"math/rand"
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
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	fmt.Println(i)
}
func main() {
	//f()

	for i := 1; i < 10; i++ {
		go f1()
	}
}
