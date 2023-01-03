package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			s += "." + suffix
		}
		return s
	}
}

func main() {
	jpgFunc := makeSuffixFunc("jpg")
	pngFunc := makeSuffixFunc("png")
	fmt.Println(jpgFunc("xxxx"))
	fmt.Println(pngFunc("xxxx"))
}
