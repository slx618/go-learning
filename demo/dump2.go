package main

import (
	"os"
	"fmt"
	"bufio"
)

func main()  {
	counts := make(map[string]int)
	file := os.Args[1:]

	if len(file) == 0 {
		countLines(os.Stdin, counts)

	} else {
		for _, arg := range file {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dump2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", line)
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}

}