package repository

import (
	"context"

	"github.com/fiuskylab/pegasus/src/proto"
)

type (
	// GRPCRepo represents the gRPC repository.
	GRPCRepo struct {
		proto.UnimplementedPegasusServer
	}
)

// NewGRPCRepo returns a new gRPC repository.
func NewGRPCRepo() *GRPCRepo {
	return &GRPCRepo{}
}

// Send - Endpoint fo
func (g *GRPCRepo) Send(ctx context.Context, req *proto.SendRequest) (*proto.SendResponse, error) {
	return nil, nil
}

// Pop -
func (g *GRPCRepo) Pop(ctx context.Context, in *proto.PopRequest) (*proto.PopResponse, error) {
	return nil, nil
}

// Consumer -
func (g *GRPCRepo) Consumer(in *proto.ConsumerRequest, srv proto.Pegasus_ConsumerServer) error {
	return nil
}

// Producer -
func (g *GRPCRepo) Producer(srv proto.Pegasus_ProducerServer) error {
	return nil
}
