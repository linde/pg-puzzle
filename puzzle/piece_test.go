package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	threeNorthPiece := NewPiece(Unspecified, North, North, North)
	threeEastPiece := NewPiece(Unspecified, East, East, East)
	squarePiece := NewPiece(Unspecified, East, South, West, North)

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

	matches := FindPieceByState(pieces, Piece6)
	assert.NotNil(matches)
	assert.Len(matches, 1)
	p6potbelly := matches[0]
	assert.Equal(p6potbelly.Rotate().steps, NewPiece(Unspecified, West, West, East, South).steps)

	matches = FindPieceByState(pieces, Piece4)
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

func TestPieceStringer(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	pieces := DefaultPieces()
	assert.NotNil(pieces)

	matches := FindPieceByState(pieces, Piece6)
	p6potbelly := matches[0]
	assert.Equal(p6potbelly.String(), "Piece{6, S S N E}")

	matches = FindPieceByState(pieces, Piece2)
	p2ell := matches[0]
	assert.Equal(p2ell.String(), "Piece{2, S E}")
}
