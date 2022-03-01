package logagent

import (
	"google.golang.org/grpc"
)

type gRPCLogAgentServer struct {
	*grpc.Server
	address string
}
