package repository

import (
	"context"
	"time"

	"github.com/fiuskylab/pegasus/src/manager"
	"github.com/fiuskylab/pegasus/src/message"
	"github.com/fiuskylab/pegasus/src/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// GRPCRepo represents the gRPC repository.
	GRPCRepo struct {
		proto.UnimplementedPegasusServer
		mgr *manager.Manager
	}
)

// NewGRPCRepo returns a new gRPC repository.
func NewGRPCRepo(mgr *manager.Manager) *GRPCRepo {
	return &GRPCRepo{
		mgr: mgr,
	}
}

// CreateTopic -
func (g *GRPCRepo) CreateTopic(ctx context.Context, req *proto.CreateTopicRequest) (*proto.CreateTopicResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	errChan := make(chan error, 1)

	go func() {
		errChan <- g.mgr.NewTopic(req.Name)
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return &proto.CreateTopicResponse{
			Error: "",
		}, nil
	}
}

// Send - Endpoint fo
func (g *GRPCRepo) Send(ctx context.Context, req *proto.SendRequest) (*proto.SendResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	msgChan := make(chan *message.Message, 1)
	errChan := make(chan error, 1)
	go func() {
		msg, err := message.FromRequest(req)
		if err != nil {
			errChan <- err
			return
		}
		if err = g.mgr.Send(msg); err != nil {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		msgChan <- msg
	}()
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		zap.L().Error(err.Error())
		return nil, status.Error(codes.Canceled, err.Error())
	case <-msgChan:
		return &proto.SendResponse{
			Message: "message sent!",
		}, nil
	}
}

// Pop -
func (g *GRPCRepo) Pop(ctx context.Context, req *proto.PopRequest) (*proto.PopResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	msgChan := make(chan *message.Message, 1)
	errChan := make(chan error, 1)

	go func() {
		msg, err := g.mgr.Pop(req.TopicName)
		if err != nil {
			errChan <- err
			return
		}
		msgChan <- msg
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		zap.L().Error(err.Error())
		return nil, status.Error(codes.Canceled, err.Error())
	case msg := <-msgChan:
		return &proto.PopResponse{
			TopicName: msg.TopicName,
			Body:      msg.Body,
		}, nil
	}
}

// Consumer -
func (g *GRPCRepo) Consumer(in *proto.ConsumerRequest, srv proto.Pegasus_ConsumerServer) error {
	return nil
}

// Producer -
func (g *GRPCRepo) Producer(srv proto.Pegasus_ProducerServer) error {
	return nil
}
