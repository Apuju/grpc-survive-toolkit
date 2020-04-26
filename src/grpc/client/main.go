package main

import (
	"context"
	"fmt"
	grpcTest "github.com/apuju/grpc-survive-toolkit/src/grpc/test_v1"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc server connects failed", err)
	}

	client := grpcTest.NewEchoServerClient(connection)
	response, err := client.Echo(context.Background(), &grpcTest.EchoRequest{
		Message: "hello, here is client",
	})

	if response != nil {
		fmt.Println("message came from server:", response.Message, ";unixtime:", response.Unixtime)
	} else {
		fmt.Print("message is null")
	}
}
