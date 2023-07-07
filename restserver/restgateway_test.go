package restserver

import (
	"bytes"
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

	tests := []struct {
		stops         []puzzle.Loc
		respCode      int
		solveExpected bool
	}{

		// TODO this always ends up as an empty array in the rpc server.

		{[]puzzle.Loc{}, http.StatusOK, true},

		// TODO there really has to be a better way to create these
		{[]puzzle.Loc{{R: 0, C: 0}, {R: 0, C: 4}, {R: 4, C: 2}}, http.StatusOK, true},

		// test some failure modes
		{[]puzzle.Loc{{R: 9999, C: 0}, {R: 0, C: 4}, {R: 4, C: 2}}, http.StatusOK, true},
	}

	for idx, test := range tests {

		testName := fmt.Sprintf("TestRestGateway(idx:%d){solveExpected:%v,stops:%v",
			idx, test.solveExpected, test.stops)

		t.Run(testName, func(tt *testing.T) {

			assert := a.New(tt)

			stopSetBytes := []byte{}

			assert.NotNil(test.stops)
			if len(test.stops) > 0 {
				// TODO do this somehow better or document it
				rpcRequest := struct {
					StopSet []puzzle.Loc
				}{
					StopSet: test.stops,
				}
				stopSetBytes, _ = json.Marshal(rpcRequest)
			}

			// TODO is the URL available in the proto generated code?
			posturl := fmt.Sprintf("http://%s/v1/puzzle/solve", gwAddr)

			resp, reqErr := http.Post(posturl, "application/json", bytes.NewBuffer(stopSetBytes))
			assert.Nil(reqErr)
			assert.NotNil(resp)
			assert.Equal(test.respCode, resp.StatusCode)
			if resp.StatusCode == http.StatusOK {
				assert.Equal(http.StatusOK, resp.StatusCode)
				defer resp.Body.Close()

				body, bodyReadErr := io.ReadAll(resp.Body)
				assert.Nil(bodyReadErr)

				bodyStr := string(body)

				var jsonResult proto.SolveReply
				jsonErr := json.Unmarshal([]byte(bodyStr), &jsonResult)

				assert.Nil(jsonErr)
				assert.Equal(test.solveExpected, jsonResult.Solved)
				if jsonResult.Solved {
					assert.Len(jsonResult.Solution, puzzle.BOARD_DIMENSION*puzzle.BOARD_DIMENSION)
				}
			}
		})
	}

}
