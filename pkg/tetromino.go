package tetris_optimizer

import (
	"reflect"
)

type (
	Board       [][]uint8
	Tetromino2D [4][4]uint8
)

type Tetromino struct {
	Id    int
	Kind  int
	Table Tetromino2D
}

type Point struct {
	X int
	Y int
}

const (
	// Square tetromino
	TT_Sq = iota

	// Bar tetromino
	TT_HB
	TT_VB

	// S tetromino
	TT_SH
	TT_SV

	// Z tetromino
	TT_ZH
	TT_ZV

	// L tetromino
	TT_L0
	TT_L90
	TT_L180
	TT_L270

	// J tetromino
	TT_J0
	TT_J90
	TT_J180
	TT_J270

	// T tetromino
	TT_T0
	TT_T90
	TT_T180
	TT_T270
)

var tetrominoes = []Tetromino2D{
	// Square tetromino
	{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},

	// Bar tetromino
	{
		{1, 1, 1, 1},
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

// Returns the kind of this tetromino piece.
func TetrominoKind(piece Tetromino2D) int {
	for index, tetro := range tetrominoes {
		if reflect.DeepEqual(tetro, piece) {
			return index
		}
	}
	return -1
}

// Return a copy of the tetromino relocated to top-left corner of its "container."
//
// ....  => ##..
// .##.  => ##..
// .##.  => ....
// ....  => ....
func TetrominoCrop(mask Tetromino2D) Tetromino2D {
	topMost, leftMost := 0, 0

foundTopMost:
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if mask[x][y] == 1 {
				topMost = x
				break foundTopMost
			}
		}
	}

foundLeftMost:
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if mask[y][x] == 1 {
				leftMost = x
				break foundLeftMost
			}
		}
	}

	replica := Tetromino2D{}
	for y := topMost; y < 4; y++ {
		for x := leftMost; x < 4; x++ {
			replica[y-topMost][x-leftMost] = mask[y][x]
		}
	}

	return replica
}
