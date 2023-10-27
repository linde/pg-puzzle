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

func Test_NoSubCommand(t *testing.T) {

	// cmd := NewRootCmd()
	// TODO figure out why i am only getting the first lines instead of all the output
	// GenericCommandRunner(t, cmd, "Usage:")

}

// TODO test usage without args
