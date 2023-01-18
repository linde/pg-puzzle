package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocIsLessThanOrEqual(t *testing.T) {

	loc0_0 := NewLoc(0, 0)
	loc0_10 := NewLoc(0, 10)
	loc10_0 := NewLoc(10, 0)

	tests := []struct {
		expected bool
		l, r     Loc
	}{
		{true, loc0_0, loc10_0},
		{true, loc0_0, loc0_10},
		{true, loc0_0, loc0_0},
		{false, loc10_0, loc0_0},
		{false, loc0_10, loc0_0},

		{true, NewLoc(0, -1), loc0_0}, // TODO should NewLoc() IsLessThanOrEqual() return an error?
	}

	for i, tt := range tests {

		testName := fmt.Sprintf("TestLocIsLessThanOrEqual_%v<=%v:%v", tt.l, tt.r, tt.expected)
		t.Run(testName, func(ttt *testing.T) {
			assertNested := assert.New(ttt)

			assertFunc := assertNested.Truef
			if !tt.expected {
				assertFunc = assertNested.Falsef
			}

			assertFunc(tt.l.IsLessThanOrEqual(tt.r), "iteration %d: expected %v for %v <= %v", i, tt.expected, tt.l, tt.r)
		})
	}

}
