package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepRotation(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	for _, step := range []Step{North, East, South, West} {
		stepRotated := step.Rotate()

		switch step {
		case North:
			assert.Equal(stepRotated, East)
		case East:
			assert.Equal(stepRotated, South)
		case South:
			assert.Equal(stepRotated, West)
		case West:
			assert.Equal(stepRotated, North)
		}
	}

}

func TestPiece(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	threeNorthPiece := NewPiece(Unknown, North, North, North)
	threeEastPiece := NewPiece(Unknown, East, East, East)
	squarePiece := NewPiece(Unknown, East, South, West, North)

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

	matches := FindPieceByState(pieces, Piece_6)
	assert.NotNil(matches)
	assert.Len(matches, 1)
	p6potbelly := matches[0]
	assert.Equal(p6potbelly.Rotate().steps, NewPiece(Unknown, West, West, East, South).steps)

	matches = FindPieceByState(pieces, Piece_4)
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

	matches := FindPieceByState(pieces, Piece_6)
	p6potbelly := matches[0]
	assert.Equal("Piece{Piece_6, South South North East}", p6potbelly.String())

	matches = FindPieceByState(pieces, Piece_2)
	p2ell := matches[0]
	assert.Equal("Piece{Piece_2, South East}", p2ell.String())
}
