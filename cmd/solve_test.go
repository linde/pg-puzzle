package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	cmd.SetArgs([]string{"--stops", "'{0 0} {4 0}'"})
	//cmd.SetArgs([]string{"--all", "true"})

	b := bytes.NewBufferString("")

	cmd.SetOut(b) // TODO figure out why this isnt  working
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	assert.NotNil(out)
	assert.Nil(err)

	assert.Contains(strings.ToLower(string(out)), "solved")

}
