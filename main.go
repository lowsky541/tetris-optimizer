package main

import (
	"fmt"
	"os"

	optimizer "tetris_optimizer/pkg"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("tetris-optimizer FILEPATH")
		os.Exit(1)
	}

	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	tetrominoes, err := optimizer.Parse(file)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	for _, tetromino := range tetrominoes {
		optimizer.DisplayTetro(tetromino.Table)
		fmt.Println()
	}

	board := optimizer.OptimizeTetrominoes(tetrominoes)

	optimizer.DisplayBoardAlpha(board)
	fmt.Println("Tetrominoes are packed up.")
}
