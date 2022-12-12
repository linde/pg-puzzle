package puzzle

// TODO consider other loop arrangements to make the logic cleaner to remove
// duplication of curVal checks and step increments

func IsSafePlacement(p *Piece, b *Board, r, c int, val State) (bool, *Board) {

	// TODO -- check to make sure the piece isnt already present on the board.

	// initialize a board we'll work with that has starts as the same as our arg
	retBoard := b.Clone()

	curR := r
	curC := c

	// first check the starting location to make sure it's unoccupied.
	// if not, mark it and keep going.
	curVal := retBoard.Get(curR, curC)
	if curVal != Empty {
		return false, nil
	}
	retBoard.Set(curR, curC, val)

	// the check the landing spot for each subsequent step to make sure is value
	// before moving to it.
	//
	// it is ok in this case for the piece val to exist because we
	// express shapes by backtracing. for example,  E E E W S S
	// makes a T shapae
	for _, step := range p.steps {

		curR, curC = doStep(curR, curC, step)
		curVal = retBoard.Get(curR, curC)

		//fmt.Printf("IsSafePlacement: From %s @ %d,%d move %s \n%s", curVal, curR, curC, step, retBoard)

		// the current value needs to be either equal to the arg val or empty, otw error
		if !(curVal == val || curVal == Empty) {
			return false, nil
		}

		// set the value for this move and check the next one
		retBoard.Set(curR, curC, val)

	}

	return true, retBoard
}
