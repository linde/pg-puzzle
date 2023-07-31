

# About
this is a a project based on a puzzle game i was given. the point is to test if the game has any configuration which is unsolvable.

<img src="doc/puzzle.jpg" width="400">

Currently, it can solve a set of stops (ie the 3 pieces you move around to make a puzzle) and also roll through and try all stops (here are the [solutions](doc/solutions.json)). This was generated by running: `go run main.go solve  --all  -o=json > doc/solutions.json`


# Build

```bash
go install golang.org/x/tools/cmd/stringer@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go get     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

go generate ./...
go build

```

# Usage

```bash
$ go run main.go solve --help
solve for default pieces

Usage:
  pg-puzzle solve [flags]

Flags:
  -a, --all            try every stop combination, not allowed with --stops
  -h, --help           help for solve
  -s, --stops string   board stops to solve, '[0-4],[0-4] [0-4],[0-4] [0-4],[0-4]' (default "0,0 0,4 4,2")
  -v, --verbose        with --all, print the solutions
  -n, --workers int    number of workers for --all (default 8)
```


# GRPC and Rest gateway

This now also has a `server` subcommand that will stand up a grpc server and also a rest gateway to proxy to it. To try out the rest gateway, do the following:

```bash

go run main.go server --rest=8080 &
curl -X 'POST'   'http://localhost:8080/v1/puzzle/solve'   \
  -H 'accept: application/json'   \
  -H 'Content-Type: application/json'  \
  -d '{"stopSet":[{"row":0,"col":0},{"row":0,"col":4},{"row":4,"col":2}]}'

```

If you're running the gateway, the app also serves an swagger schema. The [openapiv2 schema file](./proto/puzzle.swagger.json) is generated via `go generate ./...` and accessible via http for the port used above at [openapiv2.json](http://localhost:8080/openapiv2.json). We also serve the [swagger-ui](http://localhost:8080/swagger-ui/) locally on this port as well.


# Disclaimer 

I used this as an excuse to learn features of golang, lots can be improved in the code and the code style.


# TODOs

* TODO change the name of some of the `solve.go` files all over
* TODO have log levels for server stuff
