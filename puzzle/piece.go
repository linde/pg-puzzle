package shape

import (
	"fmt"
	"strings"
)

const PIECE_DIMENSION = 3

// TODO should we have a structure accessor? tests reach in directly now
type Piece struct {
	structure Board
}

// TODO should we remove padding (empty cols or rows) to make sure it
// is the most compact representation?
func NewPiece(s [][]bool) *Piece {

	//what TODO if param matrix is wider/longer than PIECE_DIMENSION

	// nomalizedStruct := initializeStructure(PIECE_DIMENSION)

	/***
	for rowIndex, row := range s {
		for colIdx, col := range row {
			if col {
				nomalizedStruct[rowIndex][colIdx] = Occupied
			}
		}
	}
	****/

	return &Piece{}
}

// TODO should this return a pointer?
func initialize(dim int) Piece {

	/***
	m := make(Piece, dim)
	for rowIdx := range m {
		row := make(Row, dim)
		m[rowIdx] = row
	}
	return m
	***/

	return Piece{}
}

// this only really works for 3x3, do assertions (or friggin generalize)
// TODO is Rotate() twice the same as flip?
func (p Piece) Rotate() *Piece {

	/**
	rotated := initializeStructure(PIECE_DIMENSION)

	for rowIdx, row := range p.structure {

		for colIdx, col := range row {

			switch rowIdx {
			case 0:
				rotated[colIdx][2] = col
			case 1:
				rotated[colIdx][1] = col
			case 2:
				rotated[colIdx][0] = col
			}
		}
	}
	**/

	return &Piece{}
}

func (p Piece) String() string {

	var b strings.Builder

	for _, row := range p.structure {
		for _, col := range row {
			fmt.Fprintf(&b, "%s ", col)
		}
		fmt.Fprintf(&b, "\n")

	}
	return b.String()
}
