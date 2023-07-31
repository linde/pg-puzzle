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
	vShapreSolved, _ := solve(vShapedBoard, pieces)
	assert.True(vShapreSolved, "Couldnt solve:\n%s\n", vShapedBoard)

}

func TestSolveAllCapped(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	ls := NewLocalSolver()

	workers, cap := 8, 2
	solved, unsolved := ls.SolveAllStops(workers, cap)
	assert.Len(unsolved, 0)
	assert.Len(solved, GetCombosForCap(cap))

	workers, cap = 1, 1
	solved, unsolved = ls.SolveAllStops(workers, cap)
	assert.Len(unsolved, 0)
	assert.Len(solved, GetCombosForCap(cap))

}
