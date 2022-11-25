package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceCoverage(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	emptyBoard := New([][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	})
	// t.Logf("\n%s", emptyBoard)

	piece := New([][]bool{
		{true},
	})
	assert.NotNil(piece)
	// t.Logf("\n%s", piece)

	assert.NotEqualValues(emptyBoard, piece)
	assert.Equal(emptyBoard.structure[0][0], Empty)

	isSafe, results := IsSafePlacement(emptyBoard, piece)
	assert.True(isSafe)
	assert.NotNil(results)
	assert.Equal(results.structure[0][0], Occupied)

	nwOnlyBoard := New([][]bool{
		{true, false, false},
		{false, false, false},
		{false, false, false},
	})
	// t.Logf("\n%s", nwOnlybnwOnlyBoardoard)

	nwIsSafe, results := IsSafePlacement(nwOnlyBoard, piece)
	assert.False(nwIsSafe)
	assert.Nil(results)

	neOnlyBoard := New([][]bool{
		{false, false, true},
		{false, false, false},
		{false, false, true},
	})
	// t.Logf("\n%s", nwOnlyboard)

	neIsSafe, results := IsSafePlacement(neOnlyBoard, piece)
	t.Logf("\n%s", results)

	assert.True(neIsSafe)
	assert.NotNil(results)
	assert.Equal(results.structure[0][0], Occupied)
	assert.Equal(results.structure[0][2], Occupied)
	assert.Equal(results.structure[2][2], Occupied)

}
