package server

import (
	"context"
	"errors"
	"fibonacci"
	"fibonacci/pkg/api"
	"fibonacci/pkg/repository"
)

type GRPCServer struct {
	repos *repository.Repository
}

func NewGRPCServer(repos *repository.Repository) *GRPCServer {
	return &GRPCServer{repos: repos}
}

func (s *GRPCServer) GetSequence(ctx context.Context, req *api.Request) (*api.Response, error) {
	if req.GetStart() < 0 || req.GetEnd() > 10000 || req.GetEnd() < req.GetStart() {
		return nil, errors.New("invalid parametrs")
	}
	input := new(fibonacci.Input)
	input.Start = req.GetStart()
	input.End = req.GetEnd()

	response := new(api.Response)

	if response.Result = s.repos.GetSequence(input); response.Result == nil {
		return nil, errors.New("internal error")
	}

	return response, nil

}
