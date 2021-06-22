package main

import (
	"fmt"
	"github.com/slx618/day05/logger"
	"io"
	"os"
	"strings"
)

func copyFile(target string) {
	handle, err := os.Open(target)
	if err != nil {
		fmt.Println(err)
	}
	defer handle.Close()
	var content []byte
	var tmp [8]byte
	for {
		n, err := handle.Read(tmp[:])
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, tmp[:n]...)

		if n < 8 {
			break
		}

	}

	//fmt.Println(string(content))

	nameList := strings.Split(target, ".go")
	s := append(nameList[:], "_copy", ".go")
	copyTarget := strings.Join(s, "")
	writeHandle, err := os.OpenFile(copyTarget, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer writeHandle.Close()
	if err != nil {
		fmt.Println(err)
	}

	writeHandle.WriteString(string(content))

}

func main() {
	//copyFile("./main.go")

	log := logger.NewLogger("xx.log")
	log.Debug("ssss")

	m1 := make(map[string]interface{}, 20)
	m1["sdsa"] = []string{"232", "323232"}
	m1["sdf"] = []interface{}{1212, "12342", []int{1323, 3, 2, 32, 3, 2, 32, 222}}
	log.Info(m1)

}
