package tetris_optimizer

import (
	"fmt"
	"io"
)

const (
	STATE_TETROMINO = iota
	STATE_SEPARATOR = iota
)

func tetrominize(lines [4]string, id uint8) (Tetromino, bool) {
	table := Tetromino2D{}

	for li, line := range lines {
		for ci, char := range line {
			if char == '#' {
				table[li][ci] = 1
			} else if char == '.' {
				table[li][ci] = 0
			}
		}
	}
	table = TetrominoCrop(table)
	kind := TetrominoKind(table)

	if kind == -1 {
		return Tetromino{}, false
	}

	return Tetromino{Table: table, Kind: kind, Id: id}, true
}

func check_tetro_line(text string) bool {
	for i, tc := range text {
		if i < 4 && !(tc == '#' || tc == '.') {
			return false
		} else if i == 4 && tc != '\n' {
			return false
		}
	}
	return true
}

func Parse(reader io.Reader) ([]Tetromino, error) {
	tetrominoes := []Tetromino{}

	state := STATE_TETROMINO // current parsing state
	tetroN := uint8(0)       // count of tetrominoes

	for {
		if state == STATE_TETROMINO {
			lines := [4]string{}
			for i := 0; i < 4; i++ {
				bytes := make([]byte, 5)
				_, err := io.ReadAtLeast(reader, bytes, 5)
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					return nil, fmt.Errorf("expected tetromino; got EOF")
				}

				if !check_tetro_line(string(bytes)) {
					return nil, fmt.Errorf("invalid tetromino definition")
				}

				// drop the newline character and insert that line
				bytes = bytes[:len(bytes)-1]
				lines[i] = string(bytes)
			}

			tetroN++

			tetromino, valid := tetrominize(lines, tetroN)
			if !valid {
				return nil, fmt.Errorf("Invalid tetromino")
			}

			tetrominoes = append(tetrominoes, tetromino)

			state = STATE_SEPARATOR
		} else if state == STATE_SEPARATOR {
			needle := make([]byte, 1)
			_, err := io.ReadAtLeast(reader, needle, 1)
			if err == io.EOF {
				break
			} else if needle[0] != '\n' {
				return nil, fmt.Errorf("expected newline")
			}
			state = STATE_TETROMINO
		}
	}

	return tetrominoes, nil
}
