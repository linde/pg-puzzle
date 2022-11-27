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
	board := initializeBoard(boardSize)
	assert.NotNil(board)
	assert.Equal(len(board), boardSize)

	for _, row := range board {
		assert.Equal(len(row), boardSize)
		for _, location := range row {
			assert.Equal(location, Empty)
		}
	}
}

func TestLocationStringer(t *testing.T) {

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

	type LocationTest struct {
		b        *Board
		expected Location
		r, c     int
	}
	tests := []LocationTest{
		{b, Occupied, 0, 0},
		{b, Occupied, 0, 1},
		{b, Empty, 0, 2},
		{b, Occupied, 1, 0},
		{b, Empty, 1, 2},
		{b, Empty, BOARD_DIMENSION - 1, BOARD_DIMENSION - 1},
	}

	for _, test := range tests {
		obs := test.b.GetLocation(test.r, test.c)
		const msgFmt = "for (%d,%d) expected %s, got %s"
		msg := fmt.Sprintf(msgFmt, test.r, test.c, test.expected, obs)
		assert.Equal(obs, test.expected, msg)

	}

}

// TODO check bad locations too, should be INVALID
// TODO test clone!
