package main

import (
	"fmt"
	echo "github.com/apuju/grpc-survive-toolkit/src/grpc"
	grpcTest "github.com/apuju/grpc-survive-toolkit/src/grpc/test_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Print("server starting...\n")

	grpcServer := grpc.NewServer()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Print("grpc serving\n")
	grpcTest.RegisterEchoServerServer(grpcServer, &echo.Echo{})
}
