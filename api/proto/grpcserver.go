package proto

/* import (
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

func (s *GRPCServer) GetSequence(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	if req.GetStart() < 0 || req.GetEnd() > 93 || req.GetEnd() < req.GetStart() {
		return nil, &rpcError{}
	}
	input := new(fibonacci.Input)
	input.Start = req.GetStart()
	input.End = req.GetEnd()

	response := new(AddResponse)

	response.Result = s.repos.GetSequence(input)

	return response, nil

}

func (s *GRPCServer) mustEmbedUnimplementedFibonacciServer() {}

type rpcError struct{}

func (e *rpcError) Error() string {
	return "invalid parametrs"
} */
