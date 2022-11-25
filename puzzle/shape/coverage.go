package shape

func IsSafePlacement(board *Piece, piece *Piece) (bool, *Piece) {

	// TODO should i be cloning the board on false
	// TODO should we assert that pieces cannot have blocked values

	// first check the case that these are empty structures
	if len(board.structure) == 0 && len(piece.structure) == 0 {
		return false, nil
	}

	returnBoard := initializeMatrix(len(board.structure))

	for boardRowIdx, boardRow := range board.structure {

		if boardRowIdx >= len(piece.structure) {
			// TODO this can be ok if the piece structure is smaller
			return false, nil
		}
		pieceRow := piece.structure[boardRowIdx]

		for boardColIdx, boardCol := range boardRow {
			if boardColIdx >= len(pieceRow) {
				// TODO this can be ok if the piece structure is smaller
				return false, nil
			}
			pieceCol := pieceRow[boardColIdx]
			if pieceCol == Occupied &&
				(boardCol == Occupied || boardCol == Blocked) {
				return false, nil
			} else {
				// set the value of the return board, default to the board's
				// but if Occupied in the piece, use that
				returnBoard[boardRowIdx][boardColIdx] = boardCol
				if pieceCol == Occupied {
					returnBoard[boardRowIdx][boardColIdx] = Occupied
				}

			}

		}
	}

	return true, &Piece{structure: returnBoard}
}
