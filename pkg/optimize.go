package tetris_optimizer

import (
	"fmt"
	"math"
	"time"
)

const (
	DYNAMIC_PRINTING_ENABLED bool          = false
	DYNAMIC_PRINTING_STEP    time.Duration = 1 * time.Millisecond
)

// Optimizer private functions

func init_board(w int, h int) Board {
	a := make(Board, h)
	for i := range a {
		a[i] = make([]uint8, w)
	}
	return a
}

func put_tetromino(board *Board, tetromino Tetromino, point Point, id uint8) bool {
	// Determine if the piece will break free of the board
	for py := 0; py < 4; py++ {
		for px := 0; px < 4; px++ {
			if tetromino.Table[py][px] != 0 {
				x, y := point.X+px, point.Y+py
				if (x < 0 || x >= len(*board)) || (y < 0 || y >= len(*board)) {
					return false
				}

				if (*board)[y][x] != 0 {
					return false
				}
			}
		}
	}

	// Copy the tetromino on the board
	for py := 0; py < 4; py++ {
		for px := 0; px < 4; px++ {
			if tetromino.Table[py][px] != 0 {
				x, y := point.X+px, point.Y+py
				(*board)[y][x] = id
			}
		}
	}

	return true
}

func rem_tetromino(board *Board, tetromino Tetromino, point Point) {
	for py := 0; py < 4; py++ {
		for px := 0; px < 4; px++ {
			if tetromino.Table[py][px] != 0 {
				x, y := point.X+px, point.Y+py
				(*board)[y][x] = 0
			}
		}
	}
}

func backtrack(board *Board, stack []Tetromino, index uint8) bool {
	if int(index) == len(stack) {
		return true
	}

	for y := 0; y < len(*board); y++ {
		for x := 0; x < len(*board); x++ {
			if put_tetromino(board, stack[index], Point{x, y}, uint8(index+1)) {

				if DYNAMIC_PRINTING_ENABLED {
					DisplayBoardAlpha(*board)
					time.Sleep(DYNAMIC_PRINTING_STEP)
					fmt.Println("\033[2J\033[0;0H")
				}

				if backtrack(board, stack, index+1) {
					return true
				}

				rem_tetromino(board, stack[index], Point{x, y})
			}
		}
	}

	return false
}

// Optimizer public functions

func OptimizeTetrominoes(tetrominoes []Tetromino) Board {
	success := false

	// Minimum board size capable of handling all the tetrominoes
	boardSize := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))

	board := init_board(boardSize, boardSize)
	for !success {
		if !backtrack(&board, tetrominoes, 0) {
			boardSize++
			board = init_board(boardSize, boardSize)
		} else {
			success = true
		}
	}

	return board
}
