package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolving(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	pieces := DefaultPieces()

	vShapedBoard := NewEmptyBoard()
	vShapedBoard.Set(Blocked, Loc{0, 0}, Loc{0, 4}, Loc{4, 2})
	vShapreSolved, _ := Solve(vShapedBoard, pieces)
	assert.True(vShapreSolved, "Couldnt solve:\n%s\n", vShapedBoard)

}
