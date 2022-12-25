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

const (
	Empty State = iota
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

type Loc struct {
	r int
	c int
}

func NewLoc(r, c int) Loc {
	return Loc{r, c}
}

type StopPair [2]Loc

func NormalizedStopPair(loc1, loc2 Loc) StopPair {

	if loc1.r < loc2.r {
		return StopPair{loc1, loc2}
	} else if loc1.r == loc2.r && loc1.c <= loc2.c {
		return StopPair{loc1, loc2}
	}
	return StopPair{loc2, loc1}
}

// TODO should this return a pointer?
func NewEmptyBoard(dim int) Board {

	board := make(Board, dim)
	for rowIdx := range board {
		row := make(Row, dim)
		board[rowIdx] = row
	}
	return board
}

func (b *Board) SetStopPair(val State, stops StopPair) {
	for _, loc := range stops {
		b.Set(loc.r, loc.c, val)
	}
}

// TODO stop pairs usually work for all these, right?
func (b *Board) SetN(val State, locs ...Loc) {
	for _, loc := range locs {
		b.Set(loc.r, loc.c, val)
	}
}

func NewBoard(s [][]bool) *Board {

	board := NewEmptyBoard(BOARD_DIMENSION)

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
	board := NewEmptyBoard(len(*orig))

	for rowIndex, row := range *orig {
		copy(board[rowIndex], row)
	}
	return &board
}

func (b Board) Get(row, col int) State {

	if row < 0 || row >= len(b) {
		return Invalid
	}
	if col < 0 || col >= len(b[row]) {
		return Invalid
	}

	return b[row][col]
}

func (b Board) Set(row, col int, val State) {

	// TODO need tests and sanity checks for the setters and getters
	b[row][col] = val
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
