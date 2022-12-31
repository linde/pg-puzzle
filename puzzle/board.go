package puzzle

import (
	"fmt"
	"strings"
)

const BOARD_DIMENSION = 5

// TODO fun exercise, port this to bitsets

type State int
type Row []State
type Board []Row

// TODO should Unspecified be the 0 value? we'd need to explicitly write empty in NewEmptyBoard
const (
	Empty State = iota
	Unspecified
	Occupied
	Blocked
	Invalid
	Piece1
	Piece2
	Piece3
	Piece4
	Piece5
	Piece6
)

// TODO if we change Unspecified to be the zero State, we need to copy Empty in the cells
func NewEmptyBoard(dim int) *Board {

	board := make(Board, dim)
	for rowIdx := range board {
		row := make(Row, dim)
		board[rowIdx] = row
	}
	return &board
}

func (b Board) SetStops(val State, stops StopSet) *Board {

	// TODO should be able to use SetN() with a stops[:]..., right?

	for _, loc := range stops {
		b.Set(loc, val)
	}
	return &b
}

// TODO stop pairs usually work for all these, right?
func (b Board) SetN(val State, locs ...Loc) {
	for _, loc := range locs {
		b.Set(loc, val)
	}
}

func NewBoard(s [][]bool) *Board {

	board := NewEmptyBoard(BOARD_DIMENSION)

	// TODO: should we assert that s < BOARD_DIMENSION?
	for rowIndex, row := range s {
		for colIdx, col := range row {
			if col {
				(*board)[rowIndex][colIdx] = Occupied
			}
		}
	}

	return board
}

func (orig Board) Clone() (neb *Board) {

	// TODO sanity check orig cols dimensions?
	neb = NewEmptyBoard(len(orig))

	for rowIdx, row := range orig {
		for colIdx, val := range row {
			neb.SetN(val, Loc{rowIdx, colIdx})
		}
	}
	return
}

func (b Board) Get(loc Loc) State {

	if loc.r < 0 || loc.r >= len(b) {
		return Invalid
	}
	if loc.c < 0 || loc.c >= len(b[loc.r]) {
		return Invalid
	}

	return b[loc.r][loc.c]
}

func (b Board) Set(loc Loc, val State) {

	// TODO need tests and sanity checks for the setters and getters
	b[loc.r][loc.c] = val
}

// TODO make a stringer that prints multiple boards side by side via ...Board
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

func ParallelBoardsString(boards ...*Board) string {

	var b strings.Builder

	maxRows := 0
	for _, board := range boards {
		if board != nil && len(*board) > maxRows {
			maxRows = len(*board)
		}
	}

	if maxRows == 0 {
		fmt.Fprintf(&b, "%d nil boards\n", len(boards))
	}

	for i := 0; i < maxRows; i++ {
		fmt.Fprintf(&b, "|")
		for _, board := range boards {
			if board == nil {
				nilBoardRow := "         "
				if i == 2 {
					// TODO prob should make sure there is a 2nd row or this wont print
					nilBoardRow = "   nil   "
				}
				fmt.Fprintf(&b, "%s", nilBoardRow)
			} else {
				delim := ""
				for _, col := range (*board)[i] {
					fmt.Fprintf(&b, "%s%s", delim, col)
					delim = " "
				}
			}
			fmt.Fprintf(&b, "|")
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

func (e State) String() string {
	switch e {
	case Empty:
		return "E"
	case Occupied:
		return "O"
	case Blocked:
		return "B"
	case Piece1:
		return "1"
	case Piece2:
		return "2"
	case Piece3:
		return "3"
	case Piece4:
		return "4"
	case Piece5:
		return "5"
	case Piece6:
		return "6"
	}
	return "?"
}
