package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	threeNorthPiece := NewPiece(North, North, North)
	threeEastPiece := NewPiece(East, East, East)
	squarePiece := NewPiece(East, South, West, North)

	assert.NotNil(threeNorthPiece)
	assert.NotNil(threeEastPiece)
	assert.NotNil(squarePiece)

	rotatedThreeNorthPiece := threeNorthPiece.Rotate()
	assert.NotEqualValues(*threeNorthPiece, *rotatedThreeNorthPiece)
	assert.EqualValues(*rotatedThreeNorthPiece, *threeEastPiece)

	// do we care that we cant tell a rotated square is the same thing?
	// rotatedSquarePiece := squarePiece.Rotate()
	// assert.EqualValues(squarePiece, rotatedSquarePiece)

	// rotate thee more times back to the home orientation
	for i := 1; i <= 3; i++ {
		rotatedThreeNorthPiece = rotatedThreeNorthPiece.Rotate()
	}
	assert.EqualValues(rotatedThreeNorthPiece, threeNorthPiece)

}
