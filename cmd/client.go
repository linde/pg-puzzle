package cmd

import (
	"context"
	"fmt"
	"pgpuzzle/grpcservice"
	"pgpuzzle/proto"
	"time"

	"github.com/spf13/cobra"
)

// TODO it'd be cool to have these instance scoped somehow so you avoid cross serverings

// TODO just fold this into the solve command
var clientCmd = NewClientCmd()
var clientRpcPort, timeoutSeconds int
var host string

func NewClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "client for a rpc or rest service",
		RunE:  doClientRun,
	}
	cmd.Flags().StringVarP(&host, "host", "s", "localhost", "rpcserver host")
	cmd.Flags().IntVarP(&clientRpcPort, "port", "p", DEFAULT_RPC_PORT, "rpcserver port")
	cmd.Flags().IntVarP(&timeoutSeconds, "timeout", "t", 0, "grpc client timeout (seconds)")

	return cmd
}

func init() {
	RootCmd.AddCommand(clientCmd)
}

func doClientRun(cmd *cobra.Command, args []string) error {

	var ctx = context.Background()
	if timeoutSeconds > 0 {
		timeout := time.Second * time.Duration(timeoutSeconds)
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	target := fmt.Sprintf("%s:%d", host, clientRpcPort)
	client, err := grpcservice.NewNetClientConn(ctx, target)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "did not connect: %v", err)
		return err
	}
	defer client.Close()

	request := &proto.SolveRequest{}
	reply, clientErr := client.Call(request)

	if clientErr != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "did not connect: %v", err)
		return err
	}

	solution := reply.GetSolution()

	fmt.Fprintf(cmd.ErrOrStderr(), "Got: %v\n", solution)

	return nil
}
