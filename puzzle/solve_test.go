package puzzle

import (
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
		Piece6: NewPiece(South, South, South, SkipNorth, East), // TODO no workie
	}

	assert.NotNil(pieces)

	board := NewEmptyBoard(BOARD_DIMENSION)
	//board.Set(Blocked, Rowcol{0, 0}, Rowcol{0, 1}, Rowcol{1, 0})

	isSolved, resultingBoard := Solve(&board, pieces)
	if isSolved {
		t.Logf("Solved!\n%s", resultingBoard)
	}

}
