package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewRootCmd()
	assert.NotNil(cmd)

	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	assert.Nil(err)
	assert.Contains(strings.ToLower(string(out)), "cli")

}
