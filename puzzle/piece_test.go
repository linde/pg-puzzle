package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	piece := NewPiece([]Step{East, East, East})

	t.Logf("got: %s", piece)

	t.Logf("got: %s", South)
	t.Logf("got: %s", South+1)
	t.Logf("got: %s", (West+1)%West)

}

/**
func TestPiece(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	piece := NewPiece([][]bool{
		{true, true, true},
		{true},
	})
	assert.NotNil(piece)
	t.Logf("\n%s", piece)

	rotated := piece.Rotate()
	t.Logf("\n%s", rotated)
	assert.NotEqualValues(piece, rotated)

	expectedRotatedShape := NewPiece([][]bool{
		{false, true, true},
		{false, false, true},
		{false, false, true},
	})
	assert.EqualValues(expectedRotatedShape, rotated)

	// rotate thee more times back to the home orientation
	for i := 1; i <= 3; i++ {
		rotated = rotated.Rotate()
		// TODO we should be able to assert inequality
		t.Logf("\n%s", rotated)
	}

	assert.EqualValues(piece, rotated)

	ellPiece := NewPiece([][]bool{
		{true},
		{true, true},
		{false, true},
	})
	t.Logf("\n%s", ellPiece)
	assert.NotNil(ellPiece)

	ellRotated := ellPiece.Rotate()
	for i := 1; i <= 3; i++ {
		assert.NotEqualValues(ellPiece, ellRotated)
		ellRotated = ellRotated.Rotate()
		t.Logf("\n%s", ellRotated)
	}

	assert.EqualValues(ellPiece, ellRotated)
	t.Logf("\n%s", ellPiece.Rotate())

}


***/
