package cmd

import (
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {

	cmd := NewRootCmd()
	GenericCommandRunner(t, cmd, "cli")

	falseFlag := "--not-a-real-arg"
	cmd.SetArgs([]string{falseFlag})
	GenericCommandRunner(t, cmd, "unknown flag: "+falseFlag)
}

// TODO test usage without args
