package shape

import (
	"fmt"
	"strings"
)

const PIECE_DIMENSION = 3

type Element int64

const (
	Empty Element = iota
	Occupied
	Block
)

func (e Element) String() string {
	switch e {
	case Empty:
		return "E"
	case Occupied:
		return "O"
	case Block:
		return "B"
	}
	return "?"
}

type Piece struct {
	structure [][]Element
}

// TODO should we remove padding (empty cols or rows) to make sure it
// is the most compact representation?
func New(s [][]bool) *Piece {

	//TODO put some sanity checks on dimensions

	nomalizedStruct := make([][]Element, PIECE_DIMENSION)

	for i := 0; i < PIECE_DIMENSION; i++ {
		row := make([]Element, PIECE_DIMENSION)
		nomalizedStruct[i] = row
	}

	for rowIndex, row := range s {
		for colIdx, col := range row {
			if col {
				nomalizedStruct[rowIndex][colIdx] = Occupied
			}
		}
	}

	return &Piece{structure: nomalizedStruct}
}

// TODO should this return a pointer?
func initializeMatrix(dim int8) [][]Element {

	m := make([][]Element, dim)
	for rowIdx := range m {
		row := make([]Element, dim)
		m[rowIdx] = row
	}
	return m
}

// this only really works for 3x3, do assertions (or friggin generalize)
// TODO is Rotate() twice the same as flip?
func (p Piece) Rotate() *Piece {

	rotated := initializeMatrix(PIECE_DIMENSION)

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

	return &Piece{structure: rotated}
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
