package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceCoverage(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	// midNorthBoard := NewBoard([][]bool{
	// 	{false, false, true, false, false},
	// })
	emptyBoard := NewBoard([][]bool{})
	assert.NotNil(emptyBoard)

	nwOnlyBoard := NewBoard([][]bool{
		{true},
	})
	threeEastPiece := NewPiece([]Step{East, East, East})
	threeSouthPiece := NewPiece([]Step{South, South, South})

	isSafe, boardAfter := IsSafePlacement(threeEastPiece, nwOnlyBoard, 0, 0)
	assert.False(isSafe)
	assert.Nil(boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeEastPiece, emptyBoard, 0, 0)
	assert.True(isSafe)
	assert.NotNil(boardAfter)
	assert.NotEqualValues(emptyBoard, boardAfter)
	// t.Logf("\n%s", boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeEastPiece, emptyBoard, 0, 1)
	assert.True(isSafe)
	assert.NotNil(boardAfter)
	assert.NotEqualValues(emptyBoard, boardAfter)
	// t.Logf("\n%s", boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeEastPiece, emptyBoard, 0, 2)
	assert.True(isSafe)
	assert.NotNil(boardAfter)
	assert.NotEqualValues(emptyBoard, boardAfter)
	// t.Logf("\n%s", boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeEastPiece, emptyBoard, 0, 3)
	assert.False(isSafe)
	assert.Nil(boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeSouthPiece, emptyBoard, 0, 0)
	assert.True(isSafe)
	assert.NotNil(boardAfter)
	assert.NotEqualValues(emptyBoard, boardAfter)
	// t.Logf("\n%s", boardAfter)

	isSafe, boardAfter = IsSafePlacement(threeSouthPiece, emptyBoard, 3, 0)
	assert.False(isSafe)
	assert.Nil(boardAfter)

}
