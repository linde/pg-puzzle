package shape

import (
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

	// TODO make a func to test this and format messages
	assert.Equal(b.GetLocation(0, 0), Occupied, "0,0")
	assert.Equal(b.GetLocation(0, 1), Occupied, "0,1")
	assert.Equal(b.GetLocation(0, 2), Empty, "0,2")
	assert.Equal(b.GetLocation(1, 0), Occupied, "1,0")
	assert.Equal(b.GetLocation(1, 2), Empty, "1,2")
	assert.Equal(b.GetLocation(BOARD_DIMENSION-1, BOARD_DIMENSION-1), Empty, "last location")

}
