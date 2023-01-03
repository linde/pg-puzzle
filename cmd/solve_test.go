package cmd

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this is the default one
	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2"})
	GenericCommandRunner(t, cmd, "solving for: [{0 0} {0 4} {4 2}]", "solved: true")

	// This one requires flipping
	cmd.SetArgs([]string{"--stops=0,3 1,2 3,2"})
	GenericCommandRunner(t, cmd, "solved: true")

	capToTest := 2 // a cap of two will have 2^^3 combos because we have 3 stop paths
	capParam := fmt.Sprintf("--cap=%d", capToTest)
	targetSolvedAssertion := fmt.Sprintf("solved: %d", int(math.Pow(float64(capToTest), 3)))
	cmd.SetArgs([]string{"--all", capParam})
	GenericCommandRunner(t, cmd, targetSolvedAssertion)

	// TODO add tests for -o=json
}

func Test_ParseStops(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		arg     string
		isError bool
	}{
		{"0,0 0,4 4,2", false},
		{"", true},
		{"0,0,0 0,4 4,2", true},
		{"0 1 2", true},
	}

	for idx, test := range tests {

		ss, error := parseStop(test.arg)
		if test.isError {
			assert.NotNil(error, "test %d: expected error for: %s, got: %v", idx, test.arg, ss)
		} else {
			assert.Nil(error, "test %d: didnt expect error for: %s", idx, test.arg)
		}
	}

}
