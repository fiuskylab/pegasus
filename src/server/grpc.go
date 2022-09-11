package server

import (
	"fmt"
	"net"

	"github.com/fiuskylab/pegasus/src/proto"
	"github.com/fiuskylab/pegasus/src/repository"
	"google.golang.org/grpc"
)

type (
	// GRPC stores all info/data necessary to establish a
	// gRPC server.
	GRPC struct {
		listener net.Listener
		repo     *repository.GRPCRepo
		srv      *grpc.Server
		opts     []grpc.ServerOption
		port     uint
	}
)

// NewGRPC returns a new GRPC connector.
func NewGRPC(port uint) *GRPC {
	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	repo := repository.NewGRPCRepo()

	return &GRPC{
		port: port,
		opts: opts,
		srv:  srv,
		repo: repo,
	}
}

// Start - starts the gRPC server.
func (g *GRPC) Start() error {
	var err error
	g.listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", g.port))
	if err != nil {
		return err
	}
	proto.RegisterPegasusServer(g.srv, proto.PegasusServer(g.repo))
	return nil
}

// Close - closes the gRPC server.
func (g *GRPC) Close() error {
	g.srv.GracefulStop()
	if err := g.listener.Close(); err != nil {
		return err
	}
	return nil
}
