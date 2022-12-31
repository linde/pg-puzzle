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

	board := NewEmptyBoard()
	assert.NotNil(board)
	assert.Equal(len(*board), BOARD_DIMENSION)

	for _, row := range *board {
		assert.Equal(len(row), BOARD_DIMENSION)
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

	bLocs := BoardToLocArray([][]bool{
		{true, true},
		{true},
	})
	b := NewEmptyBoard().Set(Occupied, bLocs...)
	assert.NotNil(b)
	assert.Equal(len(*b), BOARD_DIMENSION)

	bViaSet := NewEmptyBoard()
	bViaSet.Set(Occupied, Loc{0, 0}, Loc{0, 1}, Loc{1, 0})
	assert.EqualValues(b, bViaSet, "expected:\n%s\nnot equal:\n%s", b, bViaSet)

	tests := []struct {
		b        *Board
		expected State
		loc      Loc
	}{
		{b, Occupied, Loc{0, 0}},
		{b, Occupied, Loc{0, 1}},
		{b, Empty, Loc{0, 2}},
		{b, Occupied, Loc{1, 0}},
		{b, Empty, Loc{1, 2}},
		{b, Empty, Loc{BOARD_DIMENSION - 1, BOARD_DIMENSION - 1}},
	}

	for _, tt := range tests {
		obs := tt.b.Get(tt.loc)
		const msgFmt = "for %v expected %s, got %s"
		msg := fmt.Sprintf(msgFmt, tt.loc, tt.expected, obs)
		assert.Equal(obs, tt.expected, msg)
	}

}

func TestCloneBoard(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	neb := NewEmptyBoard()
	assert.NotNil(neb)

	nebClonePtr := neb.Clone()
	assert.NotNil(nebClonePtr)
	assert.Equal(*neb, *nebClonePtr)

	// make sure the clone doesnt have refs to the orig
	neb.Set(Occupied, NewLoc(4, 4))
	assert.NotEqual(neb, *nebClonePtr)

}

func DontTestParallelBoardPrinter(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	nwLocs := BoardToLocArray([][]bool{
		{true, true},
		{true},
	})
	northWest := NewEmptyBoard().Set(Blocked, nwLocs...)

	empty := NewEmptyBoard()

	fmt.Printf("%s", ParallelBoardsString(northWest, empty, empty))

	fmt.Printf("%s", ParallelBoardsString(nil))
	fmt.Printf("%s", ParallelBoardsString(nil, nil))

	fmt.Printf("%s", ParallelBoardsString(northWest, nil))
	fmt.Printf("%s", ParallelBoardsString(nil, northWest))

}

// TODO check bad locations too, should be INVALID
// TODO board.Set() tests
