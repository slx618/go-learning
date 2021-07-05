package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func f() {
	var c chan int //nil
	for {

		//会循环得很快
		select {
		case c := <-c: //阻塞
			fmt.Println(c)
		default:
			//time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse() //这行又忘了 麻痹

	//go tool pprof cpu.pprof
	// top3
	// list funcName
	//web
	if isCPUPprof {
		f1, _ := os.Create("./cpu.pprof")
		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}

	if isMemPprof {
		f2, _ := os.Create("./mem.pprof")
		pprof.WriteHeapProfile(f2)
		defer f2.Close()
	}

	for i := 0; i < 4; i++ {
		go f()
	}

	time.Sleep(time.Second * 20)
}
