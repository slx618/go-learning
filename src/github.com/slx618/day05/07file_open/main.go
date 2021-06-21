package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer fileObj.Close()
	//var content = make([]byte, 128)
	var content [128]byte
	for {
		n, err := fileObj.Read(content[:])
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
		}

		fmt.Print(string(content[:n]))
		if n < 128 {
			return
		}
	}

}
