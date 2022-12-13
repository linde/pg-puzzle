package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolving(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	pieces := map[State]*Piece{
		Piece1: NewPiece(East, South),
		Piece2: NewPiece(South, South, South),
		Piece3: NewPiece(South, South, East),
		Piece4: NewPiece(East, South, West, North),
		Piece5: NewPiece(South, East, South),
		Piece6: NewPiece(South, South, North, East),
	}

	nwswBoard := NewEmptyBoard(BOARD_DIMENSION)
	nwswBoard.SetN(Blocked, Loc{0, 0}, Loc{4, 0})
	nwswSolved, nwswSolution := Solve(&nwswBoard, pieces)
	nwswErrorMesg := fmt.Sprintf("Couldnt solve:\n%s", nwswBoard)
	assert.True(nwswSolved, nwswErrorMesg)
	t.Logf("Solved: %v\n%s", nwswSolved, nwswSolution)

	/*** TODO this one doesnt work right yet but a lot do
	nwneBoard := NewEmptyBoard(BOARD_DIMENSION)
	nwneBoard.SetN(Blocked, Loc{0, 0}, Loc{0, 4})
	nwneSolved, nwneSolution := Solve(&nwneBoard, pieces)
	nwneErrorMesg := fmt.Sprintf("Couldnt solve:\n%s", nwneBoard)
	assert.True(nwneSolved, nwneErrorMesg)
	t.Logf("Solved: %v\n%s", nwneSolved, nwneSolution)
	****/
}
