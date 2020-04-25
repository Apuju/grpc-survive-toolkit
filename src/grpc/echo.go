package grpc

import (
	"context"
	grpcTest "github.com/apuju/grpc-survive-toolkit/src/grpc/test_v1"
)

type Echo struct{}

func (e *Echo) Echo(_ context.Context, _ *grpcTest.EchoRequest) (reponse *grpcTest.EchoReply, exception error) {
	return
}
