package cmd

import (
	"fmt"
	"pgpuzzle/grpc_solveserver"
	"pgpuzzle/grpcservice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClientCommand(t *testing.T) {
	assert := assert.New(t)

	// First create a server the client will use
	// TODO move this to a generic util we also use from server test
	gs, err := grpcservice.NewServerFromPort(0)
	assert.Nil(err, "failed to create grpcservice")
	assert.NotNil(gs, "failed to create grpcservice")

	serverAssignedPort, err := gs.GetServicePort()
	assert.Nil(err)
	assert.Positive(serverAssignedPort)

	solveServer := grpc_solveserver.NewSolveServer()
	assert.NotNil(solveServer, "failed to create solveserver")
	defer solveServer.Stop()

	go gs.Serve(solveServer)

	// now actually test the client command

	// this are the default stops
	clientCmd := NewSolveCmd()

	clientCmd.SetArgs([]string{
		"--remote=true",
		fmt.Sprintf("--port=%d", serverAssignedPort),
	})
	out := GenericCommandRunner(t, clientCmd)

	assert.NotNil(out)

}
