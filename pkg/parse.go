package tetris_optimizer

import (
	"fmt"
	"io"
)

type parserState int8

const (
	STATE_TETROMINO parserState = iota
	STATE_SEPARATOR
)

// Converts an unparsed tetromino into a usable structure.
func tetrominize(lines [4]string, id int) (Tetromino, bool) {
	piece := Tetromino2D{}

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == '#' {
				piece[lineIndex][charIndex] = 1
			} else if char == '.' {
				piece[lineIndex][charIndex] = 0
			}
		}
	}

	piece = TetrominoCrop(piece)
	kind := TetrominoKind(piece)

	if kind == -1 {
		return Tetromino{}, false
	}

	return Tetromino{Table: piece, Kind: kind, Id: id}, true
}

func checkTetroLine(text string) bool {
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

	currentState := STATE_TETROMINO // current parsing state
	tetroCount := 0                 // count of tetrominoes

	for {
		if currentState == STATE_TETROMINO {
			lines := [4]string{}
			for i := 0; i < 4; i++ {
				bytes := make([]byte, 5)
				_, err := io.ReadAtLeast(reader, bytes, 5)
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					return nil, fmt.Errorf("expected tetromino; got EOF")
				}

				if !checkTetroLine(string(bytes)) {
					return nil, fmt.Errorf("invalid tetromino definition")
				}

				// drop the newline character and insert that line
				bytes = bytes[:len(bytes)-1]
				lines[i] = string(bytes)
			}

			tetroCount++

			tetromino, valid := tetrominize(lines, tetroCount)
			if !valid {
				return nil, fmt.Errorf("invalid tetromino")
			}

			tetrominoes = append(tetrominoes, tetromino)

			currentState = STATE_SEPARATOR
		} else if currentState == STATE_SEPARATOR {
			needle := make([]byte, 1)
			_, err := io.ReadAtLeast(reader, needle, 1)
			if err == io.EOF {
				break
			} else if needle[0] != '\n' {
				return nil, fmt.Errorf("expected newline")
			}
			currentState = STATE_TETROMINO
		}
	}

	return tetrominoes, nil
}
