package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceCoverage(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	/**

	emptyBoard := NewBoard([][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	})
	// t.Logf("\n%s", emptyBoard)

	nwPiece := NewPiece([][]bool{
		{true},
	})
	assert.NotNil(nwPiece)
	// t.Logf("\n%s", piece)

	assert.NotEqualValues(emptyBoard, nwPiece)
	//assert.Equal(emptyBoard.structure[0][0], Empty)


	isSafe, results := IsSafePlacement(emptyBoard, nwPiece)
	assert.True(isSafe)
	assert.NotNil(results)
	assert.Equal(results.structure[0][0], Occupied)

	nwOnlyBoard := New([][]bool{
		{true, false, false},
		{false, false, false},
		{false, false, false},
	})
	// t.Logf("\n%s", nwOnlybnwOnlyBoardoard)

	nwIsSafe, results := IsSafePlacement(nwOnlyBoard, nwPiece)
	assert.False(nwIsSafe)
	assert.Nil(results)

	neOnlyBoard := New([][]bool{
		{false, false, true},
		{false, false, false},
		{false, false, false},
	})
	// t.Logf("\n%s", nwOnlyboard)

	neIsSafe, results := IsSafePlacement(neOnlyBoard, nwPiece)
	// t.Logf("\n%s", results)

	assert.True(neIsSafe)
	assert.NotNil(results)
	assert.Equal(results.structure[0][0], Occupied)
	assert.Equal(results.structure[0][2], Occupied)

	northPiece := New([][]bool{
		{true, true, true},
	})

	northIsSafeNE, results := IsSafePlacement(neOnlyBoard, northPiece)
	assert.False(northIsSafeNE)
	assert.Nil(results)

	northIsSafeNW, results := IsSafePlacement(nwOnlyBoard, northPiece)
	assert.False(northIsSafeNW)
	assert.Nil(results)

	westPiece := New([][]bool{
		{true},
		{true},
		{true},
	})

	westIsSafeNW, results := IsSafePlacement(nwOnlyBoard, westPiece)
	assert.False(westIsSafeNW)
	assert.Nil(results)

	westIsSafeNE, results := IsSafePlacement(neOnlyBoard, westPiece)
	assert.True(westIsSafeNE)
	assert.NotNil(results)

	t.Logf("\n%s", neOnlyBoard)
	t.Logf("\n%s", westPiece)
	t.Logf("\n%s", results)

	***/

}
