//go:generate protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt=paths=source_relative puzzle.proto

//go:generate protoc --proto_path=.  --openapiv2_out .	  puzzle.proto

package proto
