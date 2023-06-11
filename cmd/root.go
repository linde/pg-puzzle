package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "cli for a program to playout puzzle configurations",
	}
}

// This represents the base command when called without any subcommands
var RootCmd = NewRootCmd()

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	// TODO do ENV and file config support too
	// TODO add a flag to start pprof as in
	// https://medium.com/@openmohan/profiling-in-golang-3e51c68eb6a8
}
