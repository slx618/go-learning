package main

import (
	"time"
	"fmt"
)

func main()  {

	today := time.Now().Weekday()
	Sar := time.Saturday
	fmt.Println(today)
	fmt.Println(today + 1)
	fmt.Println(today + 2)
	fmt.Println(Sar)

	switch Sar {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Printf("tommorrday")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("far more two days")
	}


}
