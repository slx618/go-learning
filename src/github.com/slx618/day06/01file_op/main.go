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

	var header []byte
	_, err = h.Read(header[:1])
	if err != nil {
		fmt.Println(err)
	}
	var end []byte
	handle, err := os.OpenFile("./new.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer handle.Close()

	handle.Write(header)
	handle.Write([]byte{'b'})

	for {
		n, err := h.Read(end[1:n])

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		handle.Write(end)

	}

}
