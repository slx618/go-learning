package chapter1

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(32322)
	fmt.Println("hello world:", rand.Intn(96))
}
