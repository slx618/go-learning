package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//buffio
	handle, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
	}

	defer handle.Close()
	reader := bufio.NewReader(handle)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(line)
	}
}
