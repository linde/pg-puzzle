package cmd

import (
	"fmt"

	"pgpuzzle/grpcservice"
	"pgpuzzle/solveserver"

	"github.com/spf13/cobra"
)

var serverCmd = NewServerCmd()
var rpcPort int

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "server for the greeter service",
		RunE:  doServerRun,
	}
	cmd.Flags().IntVarP(&rpcPort, "port", "p", DEFAULT_RPC_PORT, "rpcserver port")

	return cmd
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func doServerRun(cmd *cobra.Command, args []string) error {

	gs, err := grpcservice.NewServerFromPort(rpcPort)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "failed to create server: %v", err)
		return err
	}
	solveServer := solveserver.NewSolveServer()
	defer solveServer.Stop()

	serveErr := gs.Serve(solveServer)
	if serveErr != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "error in grpc server Serve(): %v", err)
		return err
	}

	return nil
}
