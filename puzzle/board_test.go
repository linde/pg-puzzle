package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: deal with and test for odd board sizes: 0, nil, -1, NaN, etc

func TestInitBoard(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	boardSize := 5
	board := NewEmptyBoard(boardSize)
	assert.NotNil(board)
	assert.Equal(len(board), boardSize)

	for _, row := range board {
		assert.Equal(len(row), boardSize)
		for _, state := range row {
			assert.Equal(state, Empty)
		}
	}
}

func TestStateStringer(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	assert.EqualValues(Occupied.String(), "O")
	assert.EqualValues(Empty.String(), "E")
	assert.EqualValues(Blocked.String(), "B")

}

func TestNewBoard(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	b := NewBoard([][]bool{
		{true, true},
		{true},
	})
	assert.NotNil(b)
	assert.Equal(len(*b), BOARD_DIMENSION)

	bViaSet := NewEmptyBoard(BOARD_DIMENSION)
	bViaSet.SetN(Occupied, Loc{0, 0}, Loc{0, 1}, Loc{1, 0})
	assert.EqualValues(b, &bViaSet, "expected:\n%s\nnot equal:\n%s", b, bViaSet)

	tests := []struct {
		b        *Board
		expected State
		r, c     int
	}{
		{b, Occupied, 0, 0},
		{b, Occupied, 0, 1},
		{b, Empty, 0, 2},
		{b, Occupied, 1, 0},
		{b, Empty, 1, 2},
		{b, Empty, BOARD_DIMENSION - 1, BOARD_DIMENSION - 1},
	}

	for _, tt := range tests {
		obs := tt.b.Get(tt.r, tt.c)
		const msgFmt = "for (%d,%d) expected %s, got %s"
		msg := fmt.Sprintf(msgFmt, tt.r, tt.c, tt.expected, obs)
		assert.Equal(obs, tt.expected, msg)
	}

}

func TestNormalizedStopPair(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	lowLoc := Loc{0, 0}
	midLoc := Loc{0, 4}
	highLoc := Loc{4, 2}
	assert.NotEqual(lowLoc, midLoc, highLoc)

	// TODO fix this after normalize works again.

	/**

	lowHighPair := StopSet{lowLoc, midLoc, highLoc}
	normedPair := NormalizedStopPair(highLoc, lowLoc)
	assert.Equal(lowHighPair, normedPair)

	equalPair := StopPair{highLoc, highLoc}
	normedEqualPair := NormalizedStopPair(highLoc, highLoc)
	assert.Equal(equalPair, normedEqualPair)
	***/
}

/**
func NormalizedStopPair(loc1, loc2 Loc) stopPair {

	// TODO unit test this quite a bit!
	if loc1.r < loc2.r {
		return stopPair{loc1, loc2}
	} else if loc1.r == loc2.r && loc1.c <= loc2.c {
		return stopPair{loc1, loc2}
	}
	return stopPair{loc2, loc1}
}
**/

func DontTestParallelBoardPrinter(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	northWest := NewBoard([][]bool{
		{true, true},
		{true},
	})

	empty := NewEmptyBoard(BOARD_DIMENSION)

	fmt.Printf("%s", ParallelBoardsString(northWest, &empty, &empty))

	fmt.Printf("%s", ParallelBoardsString(nil))
	fmt.Printf("%s", ParallelBoardsString(nil, nil))

	fmt.Printf("%s", ParallelBoardsString(northWest, nil))
	fmt.Printf("%s", ParallelBoardsString(nil, northWest))

}

// TODO check bad locations too, should be INVALID
// TODO test clone!
// TODO board.Set() tests
