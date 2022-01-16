package proto

import (
	"context"
	"fibonacci"
	"fibonacci/pkg/repository"
)

type GRPCServer struct {
	repos *repository.Repository
}

func NewGRPCServer(repos *repository.Repository) *GRPCServer {
	return &GRPCServer{repos: repos}
}

func (s *GRPCServer) GetSequence(ctx context.Context, req *Request) (*Response, error) {
	if req.GetStart() < 0 || req.GetEnd() > 10000 || req.GetEnd() < req.GetStart() {
		return nil, &rpcInvalidParametrs{}
	}
	input := new(fibonacci.Input)
	input.Start = req.GetStart()
	input.End = req.GetEnd()

	response := new(Response)

	if response.Result = s.repos.GetSequence(input); response.Result == nil {
		return nil, &rpcInternalError{}
	}

	return response, nil

}

func (s *GRPCServer) mustEmbedUnimplementedFibonacciServer() {}

type rpcInvalidParametrs struct{}
type rpcInternalError struct{}

func (e *rpcInvalidParametrs) Error() string {
	return "invalid parametrs"
}

func (e *rpcInternalError) Error() string {
	return "internal error"
}
