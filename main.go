package main

import (
	helloclient "playgen/client"
	helloserver "playgen/server"
)

//go:generate protoc --go_out=. --go-grpc_out=. proto/helloapis/messages.proto
//go:generate protoc --go_out=. --go-grpc_out=. proto/helloapis/services.proto
func main() {
	conn, err := helloclient.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go helloserver.Start()
	if err := helloclient.SayHello("vim-go"); err != nil {
		panic(err)
	}
}
