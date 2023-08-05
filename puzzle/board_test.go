package puzzle

import (
	"fmt"
	"strings"
	"testing"

	assertions "github.com/stretchr/testify/assert"
)

// TODO: deal with and test for odd board sizes: 0, nil, -1, NaN, etc

func TestInitBoard(t *testing.T) {

	assert := assertions.New(t)
	assert.NotNil(assert)

	board := NewEmptyBoard()
	assert.NotNil(board)
	assert.Equal(len(*board), BOARD_DIMENSION)

	for _, row := range *board {
		assert.Equal(len(row), BOARD_DIMENSION)
		for _, state := range row {
			assert.Equal(state, _Empty_)
		}
	}
}

func TestNewBoard(t *testing.T) {

	assert := assertions.New(t)
	assert.NotNil(assert)

	bLocs := BoardToLocArray([][]bool{
		{true, true},
		{true},
	})
	b := NewEmptyBoard().Set(Blocked, bLocs...)
	assert.NotNil(b)
	assert.Equal(len(*b), BOARD_DIMENSION)

	bViaSet := NewEmptyBoard()
	bViaSet.Set(Blocked, Loc{0, 0}, Loc{0, 1}, Loc{1, 0})
	assert.EqualValues(b, bViaSet, "expected:\n%s\nnot equal:\n%s", b, bViaSet)

	tests := []struct {
		b        *Board
		expected State
		loc      Loc
	}{
		{b, Blocked, Loc{0, 0}},
		{b, Blocked, Loc{0, 1}},
		{b, _Empty_, Loc{0, 2}},
		{b, Blocked, Loc{1, 0}},
		{b, _Empty_, Loc{1, 2}},
		{b, _Empty_, Loc{BOARD_DIMENSION - 1, BOARD_DIMENSION - 1}},
		{b, Invalid, Loc{-99, -99}},
		{b, Invalid, Loc{0, -99}},
		{b, Invalid, Loc{-99, 0}},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("%v", tt.loc)
		t.Run(testName, func(ttt *testing.T) {
			assertNested := assertions.New(ttt)

			obs := tt.b.Get(tt.loc)
			const msgFmt = "for %v expected %s, got %s"
			msg := fmt.Sprintf(msgFmt, tt.loc, tt.expected, obs)
			assertNested.Equal(obs, tt.expected, msg)
		})
	}

}

func TestCloneBoard(t *testing.T) {

	assert := assertions.New(t)
	assert.NotNil(assert)

	neb := NewEmptyBoard()
	assert.NotNil(neb)

	nebClonePtr := neb.Clone()
	assert.NotNil(nebClonePtr)
	assert.Equal(*neb, *nebClonePtr)

	// make sure the clone doesnt have refs to the orig
	loc4_4, _ := NewLoc(4, 4)
	neb.Set(Blocked, loc4_4)
	assert.NotEqual(neb, *nebClonePtr)

}

func TestParallelBoardPrinter(t *testing.T) {

	assert := assertions.New(t)
	assert.NotNil(assert)

	nwLocs := BoardToLocArray([][]bool{
		{true, true},
		{true},
	})
	northWest := NewEmptyBoard().Set(Blocked, nwLocs...)
	empty := NewEmptyBoard()

	replacements := map[string]string{
		"B": Blocked.String(),
		"E": _Empty_.String(),
	}
	nwEmptyEmptyMatch := ReplaceAll("|B B E E E|E E E E E|E E E E E|", replacements)
	assert.Contains(ParallelBoardsString(northWest, empty, empty), nwEmptyEmptyMatch)

	assert.Contains(ParallelBoardsString(nil), "1 nil boards")
	assert.Contains(ParallelBoardsString(nil, nil), "2 nil boards")

	leftNilMatch := ReplaceAll("|   nil   |E E E E E|", replacements)
	assert.Contains(ParallelBoardsString(nil, northWest), leftNilMatch)
	rightNilMatch := ReplaceAll("|E E E E E|   nil   |", replacements)
	assert.Contains(ParallelBoardsString(northWest, nil), rightNilMatch)

}

func TestBoardMatrixRender(t *testing.T) {

	assert := assertions.New(t)
	assert.NotNil(assert)

	empty := NewEmptyBoard()

	colSep, rowSep := " ", "\n"

	emptyBoardString := StringerMatrixJoin(*empty, colSep, rowSep)

	colDelimCount := strings.Count(emptyBoardString, colSep)
	rowDelimCount := strings.Count(emptyBoardString, rowSep)

	// row delim count is one fewer than the number of dimensions
	assert.Equal(BOARD_DIMENSION-1, rowDelimCount, "row delim count mismatch")
	// col delim count is one fewer than the number of cols times the number of rows
	assert.Equal((BOARD_DIMENSION-1)*BOARD_DIMENSION, colDelimCount, "column delim count mismatch")
}

func TestBoardFromIntArray(t *testing.T) {

	assert := assertions.New(t)
	const totalBoardStates = BOARD_DIMENSION * BOARD_DIMENSION

	empty := NewEmptyBoard()
	statesArray := make([]int32, totalBoardStates)
	emptyFromDefaultArray := BoardFromInt32Array(statesArray)
	assert.Equal(empty, emptyFromDefaultArray, "empty board does not match")

	const nwStateValue State = Blocked
	statesArray[0] = int32(nwStateValue)
	nwBlockedBoard := BoardFromInt32Array(statesArray)
	nwState := nwBlockedBoard.Get(Loc{0, 0})
	assert.Equal(nwStateValue, nwState, "northwest corner set in array but didnt match in board")
}
