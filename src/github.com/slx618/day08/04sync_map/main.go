package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}
var wg sync.WaitGroup

func main() {
	for i := 0; i < 23; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(i)
			set(key, n)

			fmt.Println(get(key))
		}(i)
	}

	wg.Wait()
}

func set(k string, v int) {
	m.Store(k, v)
}

func get(k string) interface{} {
	v, _ := m.Load(k)
	return v
}
