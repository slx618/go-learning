package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//handle, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer handle.Close()

	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))

}
