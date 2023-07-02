package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServerCommand(t *testing.T) {
	assert := assert.New(t)

	// TODO make this an ENV var
	skipBlockingServerTest := true
	if skipBlockingServerTest {
		return
	}
	cmd := NewServerCmd()
	assert.NotNil(cmd)

	// this are the default stops
	cmd.SetArgs([]string{})
	out := GenericCommandRunner(t, cmd)

	assert.NotNil(out)
}
