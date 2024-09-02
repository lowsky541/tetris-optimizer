package tetris_optimizer

import "fmt"

func DisplayTetro(tetro Tetromino2D) {
	for y := range tetro {
		for x := range tetro[y] {
			n := tetro[y][x]
			if n > 0 {
				fmt.Print("\u001b[31m")
			}
			fmt.Printf("%01d ", n)
			if n > 0 {
				fmt.Print("\u001b[0m")
			}
		}
		fmt.Println()
	}
}

func DisplayBoardAlpha(board Board) {
	for y := range board {
		for x := range board[y] {
			c := board[y][x]
			if c != 0 {
				fmt.Printf("%c ", 'A'+c-1)
			} else {
				fmt.Printf("%c ", '.')
			}
		}
		fmt.Println()
	}
}
