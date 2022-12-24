package puzzle

import (
	"fmt"
	"log"
	"reflect"
)

func SolveStopPair(stops StopPair) (bool, *Board) {

	pieces := GetGamePieces()

	boardToSolve := NewEmptyBoard(BOARD_DIMENSION)
	for _, loc := range stops {
		// TODO is there a way to pass an array as a variable length arg?
		boardToSolve.SetN(Blocked, loc)
	}

	boardSolved, resultBoard := Solve(&boardToSolve, pieces)

	return boardSolved, resultBoard

}

// TODO move to something more useful for slicing than a map, maybe an array of
// struct with both State and Piece
func Solve(board *Board, pieces map[State]*Piece) (bool, *Board) {

	if len(pieces) == 0 {
		return true, board
	}

	// first step, inialize the remaining pieces array which is a copy
	// of pieces with curLoc assigned and removed from it
	var curLoc State
	var curPiece *Piece

	remainingPieces := make(map[State]*Piece)

	curInitialized := false
	for l, p := range pieces {
		if !curInitialized {
			curLoc = l
			curPiece = p
			curInitialized = true
		} else {
			remainingPieces[l] = p
		}
	}

	for rowIdx, row := range *board {
		for colIdx, cell := range row {
			if cell == Empty {

				for rotationCount := 0; rotationCount < 3; rotationCount++ {

					isSafe, resultBoard := IsSafePlacement(curPiece, board, rowIdx, colIdx, curLoc)
					if isSafe {
						// fmt.Printf("success!\n%s", resultBoard)
						restSafe, restBoard := Solve(resultBoard, remainingPieces)
						if restSafe {
							return true, restBoard
						}
					}
					curPiece.Rotate()
				}

			}
		}
	}

	return false, nil
}

func SolveAllStops() {

	// TODO move print output to calling client
	solutions := make(map[StopPair]*Board)
	noSolutionsSet := make(map[StopPair]struct{})
	var EXISTS = struct{}{} // TODO make this a const

	// TODO make a boardVisitor that doesnt double visit
	for r1 := 0; r1 < BOARD_DIMENSION; r1++ {
		for c1 := 0; c1 < BOARD_DIMENSION; c1++ {
			for r2 := 0; r2 < BOARD_DIMENSION; r2++ {
				for c2 := 0; c2 < BOARD_DIMENSION; c2++ {

					loc1 := NewLoc(r1, c1)
					loc2 := NewLoc(r2, c2)

					// skip two stops on the same spot
					if reflect.DeepEqual(loc1, loc2) {
						continue
					}

					stops := NormalizedStopPair(loc1, loc2)

					// check to make sure we havent tried this normalized pair
					// TODO there must be a cooler way to visit the grid to avoid this

					if _, ok := solutions[stops]; ok {
						continue
					} else if _, ok := noSolutionsSet[stops]; ok {
						continue
					}

					log.Printf("solving for: %v ...", stops)
					boardSolved, resultBoard := SolveStopPair(stops)
					if boardSolved {
						solutions[stops] = resultBoard
					} else {
						noSolutionsSet[stops] = EXISTS
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
