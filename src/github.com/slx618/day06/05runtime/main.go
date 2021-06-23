package main

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	// main() skip => 0
	// f() skip => 1
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println(ok)
	}

	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {

	// main skip => 0
	f()

}
