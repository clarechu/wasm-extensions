package test

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"istio_ecosystem/wasm_extensions/grpc_logging"
	"net"
)

var addr = flag.String("addr", "localhost", "The address of the server to connect to")
var port = flag.String("port", "7575", "The port to connect to")

var conn *grpc.ClientConn

func init() {
	conn, _ = grpc.Dial(
		net.JoinHostPort(*addr, *port),
		grpc.WithInsecure(),
	)
}

func Write() error {
	logger := grpc_logging.NewLoggingServiceClient(conn)
	in := &grpc_logging.WriteLogRequest{
		LogEntries: []*grpc_logging.WriteLogRequest_LogEntry{},
	}
	_, err := logger.WriteLog(context.TODO(), in)
	return err
}
