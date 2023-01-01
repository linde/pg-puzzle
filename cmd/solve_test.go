package cmd

import (
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

	bothFlagsError := "Error: if any flags in the group [stops all] are set none of the others can be"
	cmd.SetArgs([]string{"--all", "--stops=0,3 1,2 3,2"})
	GenericCommandRunner(t, cmd, bothFlagsError)

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
