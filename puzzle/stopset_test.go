package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO new tests for NewStopSet()

func TestNormalizedStopSet(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	stopSets := []StopSet{
		NormalizedStopSet(Loc{0, 0}, Loc{0, 4}, Loc{0, 0}),
		NormalizedStopSet(Loc{10, 0}, Loc{4, 0}, Loc{2, 0}),
		NormalizedStopSet(Loc{10, 0}, Loc{4, 1}, Loc{4, 0}),
		NormalizedStopSet(Loc{0, 0}, Loc{4, 4}, Loc{20, 20}),
	}

	for _, ss := range stopSets {
		assert.True(ss[0].IsLessThanOrEqual(ss[1]) && ss[1].IsLessThanOrEqual(ss[2]))
	}

}

func TestDefaultStopPaths(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	path1, path2, path3 := DefaultStopPaths()

	// these stops arent reachable to be set in the puzzle
	unavailableStops := []Loc{{0, 2}, {2, 0}, {3, 3}}

	for _, stop := range unavailableStops {
		assert.NotContains(path1, stop)
		assert.NotContains(path2, stop)
		assert.NotContains(path3, stop)
	}
}
