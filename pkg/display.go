package tetris_optimizer

import "fmt"

func DisplayTetro(piece Tetromino2D) {
	for y := range piece {
		for x := range piece[y] {
			n := piece[y][x]
			fmt.Printf("%01d ", n)
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
