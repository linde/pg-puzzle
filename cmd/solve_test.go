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

	// https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra
	// TODO	cmd.SetArgs([]string{"solve", "--stops='{0 0} {4 0}'"})

	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	assert.Nil(err)
	assert.Contains(strings.ToLower(string(out)), "solved")

}
