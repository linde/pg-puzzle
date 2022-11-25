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
	t.Logf("\n%s", emptyBoard)

	piece := New([][]bool{
		{true},
	})
	assert.NotNil(piece)
	t.Logf("\n%s", piece)

	assert.NotEqualValues(emptyBoard, piece)

	isSafe, _ := IsSafePlacement(emptyBoard, piece)
	assert.True(isSafe)
	// assert.NotNil(results)

	nwOnly := [][]bool{
		{true, false, false},
		{false, false, false},
		{false, false, false},
	}
	nwOnlyboard := New(nwOnly)
	t.Logf("\n%s", nwOnlyboard)

	nwIsSafe, results := IsSafePlacement(nwOnlyboard, piece)
	assert.False(nwIsSafe)
	assert.Nil(results)

}
