package grpc

import (
	"context"
	"fmt"
	grpcTest "github.com/apuju/grpc-survive-toolkit/src/grpc/test_v1"
	"time"
)

type Echo struct{}

func (e *Echo) Echo(_ context.Context, request *grpcTest.EchoRequest) (response *grpcTest.EchoReply, exception error) {
	requestMessage := request.Message
	fmt.Println("message came from client:", requestMessage)
	response = &grpcTest.EchoReply{
		Message:  "roger;" + requestMessage,
		Unixtime: time.Now().Unix(),
	}
	return
}
