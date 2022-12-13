package puzzle

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
