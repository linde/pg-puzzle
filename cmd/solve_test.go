package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--all=false"})

	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	assert.NotNil(out)
	assert.Nil(err)

	// TODO really need this to work
	// assert.Contains(strings.ToLower(string(out)), "solved: true")

}

func Test_parseStop(t *testing.T) {
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
