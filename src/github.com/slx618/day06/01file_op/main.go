package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	h, err := os.Open("./main.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer h.Close()

	var header [1]byte
	_, err = h.Read(header[:])
	if err != nil {
		fmt.Println(err)
	}
	handle, err := os.OpenFile("./new.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer handle.Close()

	var end [128]byte
	handle.Write(header[:])
	handle.Write([]byte{'b'})
	//跳过第一个字节
	h.Seek(1, 0)
	for {
		_, err := h.Read(end[:])

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		handle.Write(end[:])

	}

}
