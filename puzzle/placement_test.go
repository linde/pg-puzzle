package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceCoverage(t *testing.T) {

	loc0_0, _ := NewLoc(0, 0)
	loc0_3, _ := NewLoc(0, 3)
	loc0_2, _ := NewLoc(0, 2)
	loc2_4, _ := NewLoc(2, 4)
	loc3_0, _ := NewLoc(3, 0)
	loc1_3, _ := NewLoc(1, 3)

	emptyBoard := NewEmptyBoard()
	nwOnlyBoard := NewEmptyBoard().Set(Blocked, loc0_0)
	midNorthBoard := NewEmptyBoard().Set(Blocked, loc0_2)
	midEastBoard := NewEmptyBoard().Set(Blocked, loc2_4)

	threeEastPiece := NewPiece(Unknown, East, East, East)
	threeSouthPiece := NewPiece(Unknown, South, South, South)
	threeNorthPiece := NewPiece(Unknown, North, North, North)
	// TODO test west moves  threeWestPiece := NewPiece(West, West, West)
	potBellyPiece := NewPiece(Unknown, South, South, North, East)

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
		{threeSouthPiece, emptyBoard, loc3_0, false},
		{threeEastPiece, midNorthBoard, loc0_2, false},
		{threeSouthPiece, midNorthBoard, loc0_2, false},

		{threeNorthPiece, nwOnlyBoard, loc0_3, false},

		{potBellyPiece, midEastBoard, loc1_3, false},
		{potBellyPiece, midEastBoard, loc0_3, true},
	}

	for testIdx, tt := range tests {

		testName := fmt.Sprintf("%v@%v:%v", tt.p, tt.loc, tt.expectValid)
		t.Run(testName, func(ttt *testing.T) {
			assert := assert.New(ttt)

			isSafe, boardAfter := IsSafePlacement(tt.p, tt.b, tt.loc)

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
		})
	}

}
