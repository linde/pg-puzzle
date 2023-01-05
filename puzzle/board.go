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

func NewEmptyBoard() *Board {

	board := make(Board, BOARD_DIMENSION)
	for rowIdx := range board {
		row := make(Row, BOARD_DIMENSION)
		board[rowIdx] = row
	}
	return &board
}

// this also works for StopSets slices
func (b Board) Set(val State, locs ...Loc) *Board {

	// TODO have sanity checks
	for _, loc := range locs {
		b[loc.r][loc.c] = val
	}
	return &b
}

func (orig Board) Clone() *Board {

	neb := make([]Row, len(orig))
	for i := range orig {
		neb[i] = make(Row, len(orig[i]))
		copy(neb[i], orig[i])
	}

	b := Board(neb)
	return &b
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

func (p Board) String() string {

	rowStrFunc := func(row Row) string {
		statesValuesFromRow := Map([]State(row), func(s State) string { return s.String() })
		return strings.Join(statesValuesFromRow, " ")
	}
	rowStringsFromBoard := Map([]Row(p), rowStrFunc)

	return strings.Join(rowStringsFromBoard, "\n")
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
