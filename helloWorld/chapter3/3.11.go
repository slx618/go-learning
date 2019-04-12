package main

import (
	"fmt"
	"strings"
)

func main()  {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)
	board[0][1] = "X"
	board[1][1] = "O"
	board[0][2] = "X"
	board[1][2] = "O"
	board[2][1] = "O"

	fmt.Println(board)


	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}