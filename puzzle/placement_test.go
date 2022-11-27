package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceCoverage(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	emptyBoard := NewBoard([][]bool{})
	nwOnlyBoard := NewBoard([][]bool{
		{true},
	})
	midNorthBoard := NewBoard([][]bool{
		{false, false, true, false, false},
	})

	threeEastPiece := NewPiece(East, East, East)
	threeSouthPiece := NewPiece(South, South, South)

	tests := []struct {
		p           *Piece
		b           *Board
		r           int
		c           int
		expectValid bool
	}{
		{threeEastPiece, nwOnlyBoard, 0, 0, false},
		{threeEastPiece, emptyBoard, 0, 0, true},
		{threeEastPiece, emptyBoard, 0, 3, false},
		{threeSouthPiece, nwOnlyBoard, 0, 0, false},
		{threeSouthPiece, emptyBoard, 0, 0, true},
		{threeSouthPiece, emptyBoard, 3, 0, false},
		{threeEastPiece, midNorthBoard, 0, 2, false},
		{threeSouthPiece, midNorthBoard, 0, 2, false},
	}

	// TODO put better messages in here
	for _, tt := range tests {

		isSafe, boardAfter := IsSafePlacement(tt.p, tt.b, tt.r, tt.c)

		if tt.expectValid {
			assert.True(isSafe)
			assert.NotNil(boardAfter)
			assert.NotEqualValues(tt.b, boardAfter)
		} else {
			assert.False(isSafe)
			assert.Nil(boardAfter)
		}

	}

}
