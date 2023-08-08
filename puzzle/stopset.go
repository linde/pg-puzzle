package puzzle

import (
	"errors"
	"fmt"
	"sort"
)

type StopSet [3]Loc

const (
	STOPS_FORMAT string = "%d,%d %d,%d %d,%d"
)

func NewStopSet(stops string) (StopSet, error) {

	var s1r, s1c, s2r, s2c, s3r, s3c int

	fmt.Sscanf(stops, STOPS_FORMAT, &s1r, &s1c, &s2r, &s2c, &s3r, &s3c)
	locs := []struct{ r, c int }{
		{s1r, s1c},
		{s2r, s2c},
		{s3r, s3c},
	}

	var parsedLocs [3]Loc

	for idx, l := range locs {
		loc, ok := NewLoc(l.r, l.c)
		if !ok {
			err := fmt.Errorf("invalid stop with index %d (%d,%d) in %s", idx, l.r, l.c, stops)
			return StopSet{}, err
		}
		parsedLocs[idx] = loc
	}

	// TODO is thers some cool golang idiomatic way to do this?
	if parsedLocs[0] == parsedLocs[1] ||
		parsedLocs[0] == parsedLocs[2] ||
		parsedLocs[1] == parsedLocs[2] {
		err := fmt.Errorf("duplicate stops in %s", stops)
		return StopSet{}, err
	}

	return NormalizedStopSet(parsedLocs[0], parsedLocs[1], parsedLocs[2]), nil
}

func NormalizedStopSet(loc1, loc2, loc3 Loc) StopSet {

	ss := []Loc{loc1, loc2, loc3}
	sort.Slice(ss, func(i, j int) bool { return ss[i].IsLessThanOrEqual(ss[j]) })

	// TODO is there a more idiomatic way to do this copy to an array from a slice?
	return StopSet{ss[0], ss[1], ss[2]}
}

func (ss StopSet) IsValid() (bool, error) {

	// TODO maybe accumulate the errors rather than fail fast?
	for idx, loc := range ss {

		if loc.R < 0 || loc.R > BOARD_DIMENSION || loc.C < 0 || loc.C > BOARD_DIMENSION {
			errMsg := fmt.Sprintf("Location %d in StopSet %v is invalid", idx, ss)
			return false, errors.New(errMsg)
		}
	}
	return true, nil
}

// The values below describe the stop paths for the three plugs in the puzzle.
func DefaultStopPaths() (path1, path2, path3 []Loc) {

	path1 = BoardToLocArray([][]bool{
		{true, true, false, false, false},
		{true, true, true, true, false},
		{false, true, false, false, false},
		{false, true, false, false, false},
		{false, false, false, false, false},
	})

	path2 = BoardToLocArray([][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, true, true, false},
		{true, false, true, false, false},
		{true, true, true, false, false},
	})

	path3 = BoardToLocArray([][]bool{
		{false, false, false, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, true, true},
	})

	return
}
