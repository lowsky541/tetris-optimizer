package tetris_optimizer

type (
	Board       [][]uint8
	Tetromino2D [4][4]uint8
)

type Tetromino struct {
	Id    uint8
	Kind  int
	Table Tetromino2D
}

type Point struct {
	X int
	Y int
}
