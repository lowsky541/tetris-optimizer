package main_test

import (
	"os"
	"testing"
	optimizer "tetris_optimizer/pkg"
)

type testCase struct {
	filepath string
	empties  int
}

func countEmpties(board optimizer.Board) int {
	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				count++
			}
		}
	}
	return count
}

func run(file *os.File) optimizer.Board {
	tetrominoes, err := optimizer.Parse(file)
	if err != nil {
		return nil
	}
	return optimizer.OptimizeTetrominoes(tetrominoes)
}

func runTest(tC testCase) func(t *testing.T) {
	return func(t *testing.T) {
		file, err := os.Open(tC.filepath)
		if err != nil {
			t.Fatal(err)
		}

		board := run(file)
		if tC.empties == -1 {
			if board == nil {
				return
			}
			t.Fatalf("expected %s to be invalid", tC.filepath)
		}

		empties := countEmpties(board)
		if empties != tC.empties {
			t.Fatalf("expected %s to contain %d empties, got %d", tC.filepath, tC.empties, empties)
		}
	}
}

func TestBadSamples(t *testing.T) {
	testCases := []testCase{
		{
			filepath: "samples/badformat",
			empties:  -1,
		},
		{
			filepath: "samples/badexample00",
			empties:  -1,
		},
		{
			filepath: "samples/badexample01",
			empties:  -1,
		},
		{
			filepath: "samples/badexample02",
			empties:  -1,
		},
		{
			filepath: "samples/badexample03",
			empties:  -1,
		},
		{
			filepath: "samples/badexample04",
			empties:  -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.filepath, runTest(tC))
	}
}

func TestGoodSamples(t *testing.T) {
	testCases := []testCase{
		{
			filepath: "samples/goodexample00",
			empties:  0,
		},
		{
			filepath: "samples/goodexample01",
			empties:  9,
		},
		{
			filepath: "samples/goodexample02",
			empties:  4,
		},
		{
			filepath: "samples/goodexample03",
			empties:  5,
		},
		{
			filepath: "samples/hardexam",
			empties:  1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.filepath, runTest(tC))
	}
}
