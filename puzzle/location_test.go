package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO test loc0_neg1, notOk := NewLoc(0, -1)

func TestLocs(t *testing.T) {

	values := []struct {
		r, c          int
		expectedIsOk  bool
		expectedValue *Loc
	}{
		{0, 0, true, &Loc{}},
		{BOARD_DIMENSION, 0, false, nil},
		{0, BOARD_DIMENSION, false, nil},
		{-1, 0, false, nil},
		{0, -1, false, nil},
		{3, 3, true, &Loc{3, 3}},
	}

	for _, v := range values {

		testName := fmt.Sprintf("Loc(%d<=%d) ok(%v)", v.r, v.c, v.expectedIsOk)
		t.Run(testName, func(tt *testing.T) {
			assertNested := assert.New(tt)

			loc, isOK := NewLoc(v.r, v.c)
			assertNested.Equal(isOK, v.expectedIsOk)
			if isOK && v.expectedValue != nil {
				assertNested.Equal(loc, *v.expectedValue)
			}
		})

	}

}

func TestLocIsLessThanOrEqual(t *testing.T) {

	loc_0_0, _ := NewLoc(0, 0)
	loc_0_max, _ := NewLoc(0, BOARD_DIMENSION-1)
	loc_max_0, _ := NewLoc(BOARD_DIMENSION-1, 0)

	tests := []struct {
		expected bool
		l, r     Loc
	}{
		{true, loc_0_0, loc_max_0},
		{true, loc_0_0, loc_0_max},
		{true, loc_0_0, loc_0_0},
		{false, loc_max_0, loc_0_0},
		{false, loc_0_max, loc_0_0},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("(%v<=%v)==%v", tt.l, tt.r, tt.expected)
		t.Run(testName, func(ttt *testing.T) {
			assertNested := assert.New(ttt)
			assertNested.Equal(tt.l.IsLessThanOrEqual(tt.r), tt.expected)
		})
	}

}
