package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/puzzle"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops
	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--out=text"})
	GenericCommandRunner(t, cmd, "solved: [{0 0} {0 4} {4 2}]", "b 1 4 4 b")

	// These stops require flipping to solve
	cmd.SetArgs([]string{"--stops=0,3 1,2 3,2", "--out=text"})
	GenericCommandRunner(t, cmd, "solved: [{0 3} {1 2} {3 2}]", "1 2 2 b 6")

	capToTest := 2
	capParam := fmt.Sprintf("--cap=%d", capToTest)
	cmd.SetArgs([]string{"--all", capParam, "--out=text"})
	commandOutput := GenericCommandRunner(t, cmd)
	assert.NotNil(commandOutput)

	unsolvedPuzzlesInOutput := strings.Count(strings.ToLower(commandOutput), "unsolved")
	assert.Zero(unsolvedPuzzlesInOutput)

	// verify the number of solutions for the cap
	solvedPuzzles := puzzle.GetCombosForCap(capToTest)
	solvedPuzzlesInOutput := strings.Count(strings.ToLower(commandOutput), "solved")
	assert.Equal(solvedPuzzles, solvedPuzzlesInOutput)
}

func Test_SolveCommandColor(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops
	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--out=color"})
	// TODO run the command and make assertions about color in the output
	// GenericCommandRunner(t, cmd, "solved: [{0 0} {0 4} {4 2}]", "b 1 4 4 b")
}

func Test_SolveCommandJson(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops, make sure we can unmarshall
	// something and assert correct values
	stopStr := "0,0 0,4 4,2"
	cmd.SetArgs([]string{"--stops=" + stopStr, "--out=json"})
	commandOutput := GenericCommandRunner(t, cmd)
	assert.NotZero(len(commandOutput))

	var jsonResult []puzzle.SolveResult
	err := json.Unmarshal([]byte(commandOutput), &jsonResult)
	assert.Nil(err)
	assert.Equal(len(jsonResult), 1)
	assert.True(jsonResult[0].Solved)

	// TODO get the defaultStopSet from somewhere authoritative
	testStopSet, stopSetParseError := parseStop(stopStr)
	assert.Nil(stopSetParseError)
	assert.NotNil(testStopSet)
	// assert what comes out is what we put in
	assert.Equal(testStopSet, jsonResult[0].StopSet)

	// try a run with --all and a cap and verify the length
	capToTest := 2
	capParam := fmt.Sprintf("--cap=%d", capToTest)
	cmd.SetArgs([]string{"--all=true", capParam, "--out=json"})
	allCommandOutput := GenericCommandRunner(t, cmd)
	assert.NotZero(len(allCommandOutput))

	var allCommandJsonResult []puzzle.SolveResult

	err = json.Unmarshal([]byte(allCommandOutput), &allCommandJsonResult)
	assert.Nil(err)
	assert.Equal(len(allCommandJsonResult), puzzle.GetCombosForCap(capToTest))
	for _, result := range allCommandJsonResult {
		assert.True(result.Solved)
	}

}

func Test_ParseStops(t *testing.T) {

	tests := []struct {
		arg     string
		isError bool
	}{
		{"0,0 0,4 4,2", false},
		{"", true},
		{"0,0,0 0,4 4,2", true},
		{"0 1 2", true},
		{"0 -1 2", true},
		{"0,0 0,9 4,2", true},
		{"0,0 9,0 4,2", true},
	}
	for idx, test := range tests {

		testName := fmt.Sprintf("  [arg:%s][isError: %v]", test.arg, test.isError)
		t.Run(testName, func(tt *testing.T) {
			assertNested := assert.New(tt)
			ss, error := parseStop(test.arg)
			if test.isError {
				assertNested.NotNil(error, "test %d: expected error for: %s, got: %v", idx, test.arg, ss)
			} else {
				assertNested.Nil(error, "test %d: didnt expect error for: %s", idx, test.arg)
			}
		})
	}

}
