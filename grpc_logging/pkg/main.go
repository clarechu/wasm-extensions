package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"istio_ecosystem/wasm_extensions/grpc_logging"
	"log"
	"net"
)

var logger *log.Logger
var address string

func init() {
	flag.StringVar(&address, "address", ":7575", "grpc port default 7575")
}

func main() {
	flag.Parse()
	logger = log.Default()
	server := grpc.NewServer()
	service := new(LoggingServiceServer)
	grpc_logging.RegisterLoggingServiceServer(server, service)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("net Listen error:", err)
	}
	logger.Printf("Serving gRPC on grpc://localhost%v", address)
	server.Serve(lis)
}

type LoggingServiceServer struct {
}

func (l *LoggingServiceServer) WriteLog(ctx context.Context, req *grpc_logging.WriteLogRequest) (*grpc_logging.WriteLogResponse, error) {
	logger.Printf("log entries : %d", len(req.LogEntries))
	if len(req.LogEntries) > 0 {
		entries := req.GetLogEntries()[0]
		log.Printf("path  --> :%s", entries.Path)
		log.Printf("host  --> :%s", entries.Host)
		log.Printf("DestinationWorkload  --> :%s", entries.DestinationWorkload)
		log.Printf("SourceAddress  --> :%s", entries.SourceAddress)
	}
	return &grpc_logging.WriteLogResponse{}, nil
}
