package puzzle

import (
	"sort"
)

type StopSet [3]Loc

func NormalizedStopSet(loc1, loc2, loc3 Loc) StopSet {

	ss := []Loc{loc1, loc2, loc3}
	sort.Slice(ss, func(i, j int) bool { return ss[i].IsLessThanOrEqual(ss[j]) })

	// TODO is there a more idiomatic way to do this copy to an array from a slice?
	return StopSet{ss[0], ss[1], ss[2]}
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
