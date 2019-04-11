package main

import (
	"time"
	"fmt"
)

func main()  {
	time := time.Now()
	fmt.Println(time.Hour())
	fmt.Println(time.Minute())
	fmt.Println(time.Second())
	fmt.Println(time.Nanosecond())
	fmt.Println(time.Day())

	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning!")
	case time.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening")
	}
}
