package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCommand(t *testing.T) {

	cmd := NewRootCmd()
	GenericCommandRunner(t, cmd, "cli")

}

func GenericCommandRunner(t *testing.T, cmd *cobra.Command, outputAssertions ...string) {
	assert := assert.New(t)

	assert.NotNil(cmd)

	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.Execute()
	out, err := io.ReadAll(b)
	assert.Nil(err)
	for _, oa := range outputAssertions {
		assert.Contains(strings.ToLower(string(out)), oa)
	}

}
