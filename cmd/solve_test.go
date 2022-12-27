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

	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--all=false"})

	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	assert.NotNil(out)
	assert.Nil(err)

	assert.Contains(strings.ToLower(string(out)), "solved: true")

}
