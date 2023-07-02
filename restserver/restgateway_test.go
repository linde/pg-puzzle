package restserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pgpuzzle/grpcservice"
	"pgpuzzle/proto"
	"pgpuzzle/puzzle"
	"pgpuzzle/solveserver"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestRestGateway(t *testing.T) {

	assert := a.New(t)
	assert.NotNil(assert)

	// set up a net listenter so we can test gatewaying requests to it

	gs, err := grpcservice.NewServerFromPort(0)
	assert.NotNil(gs)
	assert.Nil(err)
	serverAssignedPort, portErr := gs.GetServicePort()
	assert.Nil(portErr)
	assert.Greater(serverAssignedPort, 0)

	solveServer := solveserver.NewSolveServer()
	defer solveServer.Stop()
	go gs.Serve(solveServer)

	// now get our grpc address and set up a rest gateway proxying to it.
	//  we use a channel to receive it's address
	gsTCPAddr, addressErr := gs.GetServiceTCPAddr()
	assert.Nil(addressErr)
	rgw := NewRestGateway(0, gsTCPAddr)
	assert.NotNil(rgw)
	defer rgw.Close()

	gwAddr := rgw.GetRestGatewayAddr()
	assert.NotNil((gwAddr))
	go rgw.Serve()

	// TODO is the URL available in the proto generated code?
	url := fmt.Sprintf("http://%s/v1/puzzle/solve", gwAddr)
	resp, httpErr := http.Get(url)
	assert.Nil(httpErr)
	assert.NotNil(resp)
	assert.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		assert.Nil(err)

		bodyStr := string(body)

		var jsonResult proto.SolveReply
		jsonErr := json.Unmarshal([]byte(bodyStr), &jsonResult)

		assert.Nil(jsonErr)
		assert.True(jsonResult.Solved)
		assert.Len(jsonResult.Solution, puzzle.BOARD_DIMENSION*puzzle.BOARD_DIMENSION)
	}

}
