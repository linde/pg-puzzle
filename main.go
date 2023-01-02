package main

import (
	"net/http"
	_ "net/http/pprof"
	"pgpuzzle/cmd"
)

func main() {

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	cmd.Execute()
}
