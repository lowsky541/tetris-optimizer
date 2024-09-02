package tetris_optimizer

import (
	"reflect"
)

const (
	// Square tetromino
	TT_Sq = iota

	// Bar tetromino
	TT_HB = iota
	TT_VB = iota

	// S tetromino
	TT_SH = iota
	TT_SV = iota

	// Z tetromino
	TT_ZH = iota
	TT_ZV = iota

	// L tetromino
	TT_L0   = iota
	TT_L90  = iota
	TT_L180 = iota
	TT_L270 = iota

	// J tetromino
	TT_J0   = iota
	TT_J90  = iota
	TT_J180 = iota
	TT_J270 = iota

	// T tetromino
	TT_T0   = iota
	TT_T90  = iota
	TT_T180 = iota
	TT_T270 = iota
)

// rules :
//  - 2 cells with minimum adjacents = 1
//  - 1 cell with minimum adjacents = 2

var Tetrominoes = []Tetromino2D{
	// Square tetromino
	{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},

	// Bar tetromino
	{
		{1, 1, 1, 1}, // 2 cells with minimum adjacents = 1
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	},

	// S tetromino
	{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	},

	// Z tetromino
	{
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 0, 0},
		{1, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
	},

	// L tetromino
	{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 1, 1, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 0, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},

	// J tetromino
	{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 0, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 1, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},

	// T tetromino
	{
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
	},
}

func TetrominoKind(piece Tetromino2D) int {
	for index, tetro := range Tetrominoes {
		if reflect.DeepEqual(tetro, piece) {
			return index
		}
	}
	return -1
}

func TetrominoOrigin(piece Tetromino2D) Point {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if piece[y][x] == 1 {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{X: -1, Y: -1}
}

func TetrominoCrop(mask Tetromino2D) Tetromino2D {
	topMost, rightMost := 0, 0

foundTopMost:
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if mask[y][x] == 1 {
				topMost = y
				break foundTopMost
			}
		}
	}

foundRightMost:
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if mask[y][x] == 1 {
				rightMost = x
				break foundRightMost
			}
		}
	}

	replica := Tetromino2D{}
	for y := topMost; y < 4; y++ {
		for x := rightMost; x < 4; x++ {
			replica[y-topMost][x-rightMost] = mask[y][x]
		}
	}

	return replica
}
