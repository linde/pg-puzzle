package cmd

import (
	"fmt"

	"pgpuzzle/grpc_solveserver"
	"pgpuzzle/grpcservice"
	"pgpuzzle/rest_solveserver"

	"github.com/spf13/cobra"
)

var serverCmd = NewServerCmd()
var restPort, rpcPort int

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "server for the greeter service",
		RunE:  doServerRun,
	}
	cmd.Flags().IntVarP(&rpcPort, "port", "p", DEFAULT_RPC_PORT, "rpcserver port")
	cmd.Flags().IntVarP(&restPort, "rest", "r", -1, "rest server port, dont set or use -1 to disable")

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
	solveServer := grpc_solveserver.NewSolveServer()
	defer solveServer.Stop()

	// if restPort is configured, add a rest gateway using the port
	if restPort >= 0 {
		rpcAddr, err := gs.GetServiceTCPAddr()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error getting RPC service address: %s", err)
			return err
		}
		rgw := rest_solveserver.NewRestGateway(restPort, rpcAddr)
		go rgw.Serve()
	}

	serveErr := gs.Serve(solveServer)
	if serveErr != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "error in grpc server Serve(): %v", serveErr)
		return serveErr
	}

	return nil
}
