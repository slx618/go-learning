package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var sourceMap = make(map[string]int, 200)

	for i := 1; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", rand.Intn(100))
		val := rand.Intn(100)
		sourceMap[key] = val
	}

	fmt.Println(sourceMap)
	var keys = make([]string, 0, 200)
	for k := range sourceMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, sourceMap[v])
	}

}
