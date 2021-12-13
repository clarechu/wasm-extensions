//go:generate protoc -I=. --proto_path=./  --proto_path=${GOPATH}/src --go_out=plugins=grpc,paths=source_relative:. log.proto

package grpc_logging
