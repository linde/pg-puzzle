package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPlayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play",
		Short: "play the game in terminal mode",
		RunE:  doPlayRun,
	}

	// these are here for tests vs in init

	// TODO validate output formats

	return cmd
}

var playCmd = NewPlayCmd()

func init() {
	RootCmd.AddCommand(playCmd)
}

func doPlayRun(cmd *cobra.Command, args []string) error {

	fmt.Printf("shall we play a game?\n")

	return nil

}
