package puzzle

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

	// TODO add tests with skip directions too

	assert.NotNil(threeNorthPiece)
	assert.NotNil(threeEastPiece)
	assert.NotNil(squarePiece)

	rotatedThreeNorthPiece := threeNorthPiece.Rotate()
	assert.NotEqualValues(*threeNorthPiece, *rotatedThreeNorthPiece)
	assert.EqualValues(*rotatedThreeNorthPiece, *threeEastPiece)

}

func TestPieceRotation(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	pieces := DefaultPieces()
	assert.NotNil(pieces)

	matches := PieceForState(pieces, Piece6)
	assert.NotNil(matches)
	assert.Len(matches, 1)
	p6potbelly := matches[0]
	assert.Equal(p6potbelly.Rotate().steps, NewPiece(West, West, East, South).steps)

	matches = PieceForState(pieces, Piece4)
	assert.NotNil(matches)
	assert.Len(matches, 1)
	p4square := matches[0]
	p4squareRotated := p4square.Rotate()
	assert.Equal(len(p4squareRotated.steps), len(p4squareRotated.steps))

	for _, step := range []Step{North, South, East, West} {
		assert.Contains(p4square.steps, step)
		assert.Contains(p4squareRotated.steps, step)
	}

	assert.EqualValues(p6potbelly.steps, p6potbelly.Rotate().Rotate().Rotate().Rotate().steps)
	assert.EqualValues(p4square.steps, p4square.Rotate().Rotate().Rotate().Rotate().steps)

}
