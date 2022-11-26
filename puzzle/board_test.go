package shape

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

	// TODO might be more go idiomatic as a table
	var checkLocation = func(b *Board, expected Location, r, c int) {
		obs := b.GetLocation(r, c)
		msg := fmt.Sprintf("for (%d,%d) expected %s, got %s", r, c, expected, obs)
		assert.Equal(obs, expected, msg)
	}

	checkLocation(b, Occupied, 0, 0)
	checkLocation(b, Occupied, 0, 1)
	checkLocation(b, Empty, 0, 2)
	checkLocation(b, Occupied, 1, 0)
	checkLocation(b, Empty, 1, 2)
	checkLocation(b, Empty, BOARD_DIMENSION-1, BOARD_DIMENSION-1)

}
