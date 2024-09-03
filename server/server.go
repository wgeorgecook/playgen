package helloserver

import (
	"log"
	"net"

	"google.golang.org/grpc"

	helloapis "playgen/gen/helloapis/v1"
)

type Server struct {
	helloapis.UnimplementedHelloServiceServer
}

func Start() {
	lis, err := net.Listen("tcp", "localhost:8008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	helloapis.RegisterHelloServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
