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
	Invalid
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

func (orig *Board) Clone() *Board {

	// TODO sanity check orig cols dimensions?
	board := initializeBoard(len(*orig))

	for rowIndex, row := range *orig {
		copy(board[rowIndex], row)
	}
	return &board
}

func (b Board) GetLocation(row, col int) Location {

	if row < 0 || row >= len(b) {
		return Invalid
	}
	if col < 0 || col >= len(b[row]) {
		return Invalid
	}

	return b[row][col]
}

func (b Board) SetLocation(row, col int, val Location) {

	// TODO need tests for the setters and getters
	b[row][col] = val
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
