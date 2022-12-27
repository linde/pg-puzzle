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
