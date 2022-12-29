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
	midEastBoard.Set(NewLoc(2, 4), Blocked)

	threeEastPiece := NewPiece(East, East, East)
	threeSouthPiece := NewPiece(South, South, South)
	threeNorthPiece := NewPiece(North, North, North)
	// TODO test west moves  threeWestPiece := NewPiece(West, West, West)
	potBellyPiece := NewPiece(South, South, North, East)

	loc0_0 := NewLoc(0, 0)
	loc0_3 := NewLoc(0, 3)
	loc0_2 := NewLoc(0, 2)

	tests := []struct {
		p           *Piece
		b           *Board
		loc         Loc
		expectValid bool
	}{

		{threeEastPiece, nwOnlyBoard, loc0_0, false},
		{threeEastPiece, emptyBoard, loc0_0, true},
		{threeEastPiece, emptyBoard, loc0_3, false},
		{threeSouthPiece, nwOnlyBoard, loc0_0, false},
		{threeSouthPiece, emptyBoard, loc0_0, true},
		{threeSouthPiece, emptyBoard, NewLoc(3, 0), false},
		{threeEastPiece, midNorthBoard, loc0_2, false},
		{threeSouthPiece, midNorthBoard, loc0_2, false},

		{threeNorthPiece, nwOnlyBoard, loc0_3, false},

		{potBellyPiece, &midEastBoard, NewLoc(1, 3), false},
		{potBellyPiece, &midEastBoard, loc0_3, true},
	}

	for testIdx, tt := range tests {

		isSafe, boardAfter := IsSafePlacement(tt.p, tt.b, tt.loc, Piece1)

		errorMsg := fmt.Sprintf("Test index: %d\n%v @ %v\nisSafe: %v\n  Before  |  After   \n%s",
			testIdx, tt.p, tt.loc, isSafe, ParallelBoardsString(tt.b, boardAfter))

		// fmt.Println(errorMsg)

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
