package helloclient

import (
	"context"
	"fmt"
	helloapis "playgen/gen/helloapis/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var helloClient helloapis.HelloServiceClient

func Connect() (*grpc.ClientConn, error) {
	if helloClient != nil {
		return nil, nil
	}
	conn, err := grpc.NewClient("localhost:8008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	helloClient = helloapis.NewHelloServiceClient(conn)

	return conn, nil
}

func SayHello(name string) error {
	req := &helloapis.HelloRequest{Name: name}
	resp, err := helloClient.Hello(context.Background(), req)
	if err != nil {
		return err
	}

	fmt.Println(resp.Greeting)
	return nil
}
