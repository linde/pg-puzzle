//go:generate stringer --type=State

package puzzle

import (
	"fmt"
	"strings"
)

const BOARD_DIMENSION = 5

// TODO fun exercise, port this to bitsets

type State int32
type Board [][]State

// TODO should Unspecified be the 0 value? we'd need to explicitly write empty in NewEmptyBoard
const (
	_Empty_ State = iota
	Unknown
	Blocked
	Invalid
	Piece_1
	Piece_2
	Piece_3
	Piece_4
	Piece_5
	Piece_6
)

func NewEmptyBoard() *Board {

	board := make(Board, BOARD_DIMENSION)
	for rowIdx := range board {
		row := make([]State, BOARD_DIMENSION)
		board[rowIdx] = row
	}
	return &board
}

// this also works for StopSets slices
func (b Board) Set(val State, locs ...Loc) *Board {

	// TODO have sanity checks
	for _, loc := range locs {
		r, c := loc.GetRowCol()
		b[r][c] = val
	}
	return &b
}

func (orig Board) Clone() *Board {

	neb := make([][]State, len(orig))
	for i := range orig {
		neb[i] = make([]State, len(orig[i]))
		copy(neb[i], orig[i])
	}

	b := Board(neb)
	return &b
}

func (b Board) Get(loc Loc) State {

	r, c := loc.GetRowCol()

	if r < 0 || r >= len(b) {
		return Invalid
	}
	if c < 0 || c >= len(b[r]) {
		return Invalid
	}

	return b[r][c]
}

func (board Board) String() string {

	replacements := map[string]string{
		Blocked.String(): "B",
		_Empty_.String(): "E",
		"Piece_":         "",
		Unknown.String(): "U",
		Invalid.String(): "X",
	}
	verboseStrings := StringerMatrixJoin(board, " ", "\n")
	replacedStr := ReplaceAll(verboseStrings, replacements)
	return replacedStr
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

func BoardFromInt32Array(ia []int32) *Board {

	b := NewEmptyBoard()

	// TODO find some cooler way to turn the array to a matrix
	for row := 0; row < BOARD_DIMENSION; row++ {
		for col := 0; col < BOARD_DIMENSION; col++ {
			idx := (row * BOARD_DIMENSION) + col
			b.Set(State(ia[idx]), Loc{R: row, C: col})
		}
	}
	return b
}
