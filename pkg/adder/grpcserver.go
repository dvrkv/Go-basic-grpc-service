package adder

import (
	"context"

	"github.com/dvrkv/grpcadder/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
