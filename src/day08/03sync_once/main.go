package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for {
		once.Do(func() {
			fmt.Println("sssss")
		})
	}
}
