package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolving(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	pieces := GetGamePieces()

	nwswBoard := NewEmptyBoard(BOARD_DIMENSION)
	nwswBoard.SetN(Blocked, Loc{0, 0}, Loc{4, 0})
	nwswSolved, _ := Solve(&nwswBoard, pieces)
	nwswErrorMesg := fmt.Sprintf("Couldnt solve:\n%s", nwswBoard)
	assert.True(nwswSolved, nwswErrorMesg)

	nwneBoard := NewEmptyBoard(BOARD_DIMENSION)
	nwneBoard.SetN(Blocked, Loc{0, 0}, Loc{0, 4})
	nwneSolved, _ := Solve(&nwneBoard, pieces)
	assert.Falsef(nwneSolved, "Shouldnt have been able to solve:\n%s", nwneBoard)
}

// TODO test SolveLocations()
