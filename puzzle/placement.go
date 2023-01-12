package puzzle

// TODO consider other loop arrangements to make the logic cleaner to remove
// duplication of curVal checks and step increments

func IsSafePlacement(p *Piece, b *Board, loc Loc) (bool, *Board) {

	// TODO -- check to make sure the piece isnt already present on the board elsewhere

	// initialize a board we'll use to follow steps and which we'll return, if placement is safe
	retBoard := b.Clone()

	curLoc := loc

	// first check the starting location to make sure it's unoccupied.
	// if not, mark it and keep going.
	curVal := retBoard.Get(curLoc)
	if curVal != _Empty_ {
		return false, nil
	}
	retBoard.Set(p.state, curLoc)

	// now check each subsequent step's loc to make sure is value
	// is Empty or our own piece's state before moving to it.
	//
	// it is ok for it to be our piece's state because some shapes
	// backtrack. for example,  {E E E W S S} makes a T shape
	for _, step := range p.steps {

		curLoc = doStep(curLoc, step)
		curVal = retBoard.Get(curLoc)

		//fmt.Printf("IsSafePlacement: From %s @ %d,%d move %s \n%s", curVal, curR, curC, step, retBoard)

		// the current value needs to be either equal to the arg val or empty, otw error
		if !(curVal == p.state || curVal == _Empty_) {
			// fmt.Printf("IsSafePlacement false: Piece %v @ %v, curLoc: %v\n%v\n", p, loc, curLoc, b)
			return false, nil
		}

		// set the value for this move and check the next one
		retBoard.Set(p.state, curLoc)

	}

	return true, retBoard
}
