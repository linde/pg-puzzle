package puzzle

import (
	"fmt"
	"log"
	"reflect"
)

func SolveStopSet(stops StopSet) (bool, *Board) {

	pieces := DefaultPieces()

	boardToSolve := NewEmptyBoard(BOARD_DIMENSION)
	boardToSolve.SetStops(Blocked, stops)
	boardSolved, resultBoard := Solve(&boardToSolve, pieces)

	return boardSolved, resultBoard

}

// TODO move to something more useful for slicing than a map, maybe an array of
// struct with both State and Piece
func Solve(board *Board, pieces []Piece) (bool, *Board) {

	if len(pieces) == 0 {
		return true, board
	}

	curPiece := &(pieces[0])
	for rotationCount := 0; rotationCount < 4; rotationCount++ {
		for rowIdx, row := range *board {
			for colIdx, cell := range row {
				if cell == Empty {
					isSafe, resultBoard := IsSafePlacement(curPiece, board, Loc{rowIdx, colIdx})
					if isSafe {
						// fmt.Printf("success!\n%s", resultBoard)
						restSafe, restBoard := Solve(resultBoard, pieces[1:])
						if restSafe {
							return true, restBoard
						}
					}
				}
			}
		}
		curPiece = curPiece.Rotate()
	}

	return false, nil
}

func SolveAllStops() {

	// TODO move print output to calling
	solutions := make(map[StopSet]*Board)
	noSolutionsSet := make(map[StopSet]struct{})
	var EXISTS = struct{}{} // TODO make this a const

	// TODO make a boardVisitor that doesnt double visit
	for r1 := 0; r1 < BOARD_DIMENSION; r1++ {
		for c1 := 0; c1 < BOARD_DIMENSION; c1++ {
			for r2 := 0; r2 < BOARD_DIMENSION; r2++ {
				for c2 := 0; c2 < BOARD_DIMENSION; c2++ {
					for r3 := 0; r3 < BOARD_DIMENSION; r3++ {
						for c3 := 0; c3 < BOARD_DIMENSION; c3++ {

							loc1 := NewLoc(r1, c1)
							loc2 := NewLoc(r2, c2)
							loc3 := NewLoc(r3, c3)

							// skip two stops on the same spot
							if reflect.DeepEqual(loc1, loc2) || reflect.DeepEqual(loc1, loc3) || reflect.DeepEqual(loc2, loc3) {
								continue
							}

							stops := NormalizedStopSet(loc1, loc2, loc3)

							// check to make sure we havent tried this normalized pair
							// TODO there must be a cooler way to visit the grid to avoid this

							if _, ok := solutions[stops]; ok {
								continue
							} else if _, ok := noSolutionsSet[stops]; ok {
								continue
							}

							log.Printf("solving for: %v ...", stops)
							boardSolved, resultBoard := SolveStopSet(stops)
							if boardSolved {
								solutions[stops] = resultBoard
							} else {
								noSolutionsSet[stops] = EXISTS
							}

						}
					}
				}
			}
		}
	}

	for locs, board := range solutions {
		fmt.Printf("Solved: %v\n%s", locs, board)
	}
	fmt.Printf("Unsolvable stops: %d\n", len(noSolutionsSet))

}
