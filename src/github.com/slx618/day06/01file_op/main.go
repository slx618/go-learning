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
	handle.Close()
	h.Close()
	err = os.Remove("./main.txt")
	fmt.Println(err)
	err = os.Rename("./new.txt", "./main.txt")
	fmt.Println(err)

}
