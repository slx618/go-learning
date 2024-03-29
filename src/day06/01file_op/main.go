package main

import (
	"fmt"
	"io"
	"os"
)

//找到main.txt 在第一个字和第二个字中间插入一个 b
func main() {

	h, err := os.Open("./main.txt")
	if err != nil {
		fmt.Println(err)
	}

	var header [1]byte
	_, err = h.Read(header[:])
	if err != nil {
		fmt.Println(err)
	}
	handle, err := os.OpenFile("./new.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	var end [128]byte
	handle.Write(header[:])
	handle.Write([]byte{'b'})
	for {
		n, err := h.Read(end[:])

		if err == io.EOF {
			handle.Write(end[:n])
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		handle.Write(end[:n])

	}
	handle.Close()
	h.Close()
	err = os.Rename("./new.txt", "./main.txt")

}
