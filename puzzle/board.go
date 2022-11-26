package shape

import (
	"fmt"
	"strings"
)

const BOARD_DIMENSION = 5

// TODO fun exercise, port this to bitsets

type Location int
type Row []Location
type Board []Row

const (
	Empty Location = iota
	Occupied
	Blocked
)

// TODO should this return a pointer?
func initializeBoard(dim int) Board {

	board := make(Board, dim)
	for rowIdx := range board {
		row := make(Row, dim)
		board[rowIdx] = row
	}
	return board
}

func NewBoard(s [][]bool) *Board {

	board := initializeBoard(BOARD_DIMENSION)

	// TODO: should we assert that s < BOARD_DIMENSION?
	for rowIndex, row := range s {
		for colIdx, col := range row {
			if col {
				board[rowIndex][colIdx] = Occupied
			}
		}
	}

	return &board
}

func (b Board) GetLocation(row, col int) Location {
	// TODO sanity checking maybs?
	return b[row][col]
}

func (p Board) String() string {

	var b strings.Builder

	for _, row := range p {
		for _, col := range row {
			fmt.Fprintf(&b, "%s ", col)
		}
		fmt.Fprintf(&b, "\n")

	}
	return b.String()
}

func (e Location) String() string {
	switch e {
	case Empty:
		return "E"
	case Occupied:
		return "O"
	case Blocked:
		return "B"
	}
	return "?"
}
