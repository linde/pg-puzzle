package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/grpc_solveserver"
	"pgpuzzle/grpcservice"
	"pgpuzzle/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClientCommand(t *testing.T) {
	assert := assert.New(t)

	// First create a server the client will use
	serverAssignedPort, serverStopFunc := setupGrpcServer(assert)
	defer serverStopFunc()

	// now actually test the client command

	clientCmd := NewSolveCmd()
	clientCmd.SetArgs([]string{
		"--remote=true",
		fmt.Sprintf("--port=%d", serverAssignedPort),
		"--output=json",
	})
	commandOutput := GenericCommandRunner(t, clientCmd)
	var jsonResult []puzzle.SolveResult
	jsonErr := json.Unmarshal([]byte(commandOutput), &jsonResult)
	assert.Nil(jsonErr)
	assert.Equal(len(jsonResult), 1)
	assert.True(jsonResult[0].Solved)

	assert.NotNil(commandOutput)

}

// this creates a gRPC server for subsequent tests which is run in its own gorountine
// it doesnt return errors back to callers because it asserts/fails if things arent
// as expected. it does return the serverAssignedPort and the stop function to defer calling
// to shut things down
func setupGrpcServer(assert *assert.Assertions) (int, func()) {

	gs, err := grpcservice.NewServerFromPort(0)
	assert.Nil(err, "failed to create grpcservice")
	assert.NotNil(gs, "failed to create grpcservice")

	serverAssignedPort, err := gs.GetServicePort()
	assert.Nil(err)
	assert.Positive(serverAssignedPort)

	solveServer := grpc_solveserver.NewSolveServer()
	assert.NotNil(solveServer, "failed to create solveserver")

	go gs.Serve(solveServer)
	return serverAssignedPort, solveServer.Stop
}
