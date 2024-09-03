package helloserver

import (
	"context"
	helloapis "playgen/gen/helloapis/v1"
)

func (s Server) Hello(ctx context.Context, in *helloapis.HelloRequest) (*helloapis.HelloResponse, error) {
	return &helloapis.HelloResponse{
		Greeting: "Hello, " + in.Name + "!",
	}, nil
}
