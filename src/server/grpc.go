package server

import (
	"fmt"
	"net"

	"github.com/fiuskylab/pegasus/src/manager"
	"github.com/fiuskylab/pegasus/src/proto"
	"github.com/fiuskylab/pegasus/src/repository"
	"google.golang.org/grpc"
)

type (
	// GRPC stores all info/data necessary to establish a
	// gRPC server.
	GRPC struct {
		listener net.Listener
		mgr      *manager.Manager
		repo     *repository.GRPCRepo
		srv      *grpc.Server
		opts     []grpc.ServerOption
	}
)

// NewGRPC returns a new GRPC connector.
func NewGRPC(mgr *manager.Manager) *GRPC {
	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	repo := repository.NewGRPCRepo()

	return &GRPC{
		mgr:  mgr,
		opts: opts,
		repo: repo,
		srv:  srv,
	}
}

// Start - starts the gRPC server.
func (g *GRPC) Start(port uint) error {
	var err error
	g.listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
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
