package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--all=false"})

	// TODO this should really solve: true
	GenericCommandRunner(t, cmd, "solving for: [{0 0} {0 4} {4 2}]", "solved: false")

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
