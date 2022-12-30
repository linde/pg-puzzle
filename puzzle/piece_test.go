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

	pieces := GetGamePieces()
	assert.NotNil(pieces)

	assert.Contains(pieces, Piece6)
	p6potbelly := pieces[Piece6]
	assert.Equal(p6potbelly.Rotate(), NewPiece(West, West, East, South))

	assert.Contains(pieces, Piece4)

	p4square := pieces[Piece4]
	p4squareRotated := p4square.Rotate()
	assert.Equal(len(p4squareRotated.steps), len(p4squareRotated.steps))

	for _, step := range []Step{North, South, East, West} {
		assert.Contains(p4square.steps, step)
		assert.Contains(p4squareRotated.steps, step)
	}

	assert.EqualValues(p6potbelly, p6potbelly.Rotate().Rotate().Rotate().Rotate())
	assert.EqualValues(p4square, p4square.Rotate().Rotate().Rotate().Rotate())

}
