package puzzle

import (
	"fmt"
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
	midEastBoard := NewEmptyBoard(BOARD_DIMENSION)
	midEastBoard.Set(2, 4, Blocked)

	threeEastPiece := NewPiece(East, East, East)
	threeSouthPiece := NewPiece(South, South, South)

	//potBellyPiece := NewPiece(South, South, South, SkipNorth, East)

	tests := []struct {
		p           *Piece
		b           *Board
		r           int
		c           int
		expectValid bool
	}{
		//{potBellyPiece, &midEastBoard, 1, 3, false},
		{threeEastPiece, nwOnlyBoard, 0, 0, false},
		{threeEastPiece, emptyBoard, 0, 0, true},
		{threeEastPiece, emptyBoard, 0, 3, false},
		{threeSouthPiece, nwOnlyBoard, 0, 0, false},
		{threeSouthPiece, emptyBoard, 0, 0, true},
		{threeSouthPiece, emptyBoard, 3, 0, false},
		{threeEastPiece, midNorthBoard, 0, 2, false},
		{threeSouthPiece, midNorthBoard, 0, 2, false},
		// {potBellyPiece, &midEastBoard, 0, 3, false},
	}

	// TODO put better messages in here
	for _, tt := range tests {

		isSafe, boardAfter := IsSafePlacement(tt.p, tt.b, tt.r, tt.c, Occupied)

		errorMsg := fmt.Sprintf("%v @ %d,%d\nisSafe: %v\n  Before  |  After   \n%s",
			tt.p, tt.r, tt.c, isSafe, ParallelBoardsString(tt.b, boardAfter))

		if tt.expectValid {
			assert.True(isSafe, errorMsg)
			assert.NotNil(boardAfter, errorMsg)
			assert.NotEqualValues(tt.b, boardAfter, errorMsg)
		} else {
			assert.False(isSafe, errorMsg)
			assert.Nil(boardAfter, errorMsg)
		}

	}

}
