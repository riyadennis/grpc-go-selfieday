package main

import (
	"context"
	"fmt"

	"github.com/grpc-go-selfieday/registerpb"
)

type server struct{}

// Registeration implement RegisterServiceServer
func (s *server) Registeration(ctx context.Context, req *registerpb.RegisterRequest) (*registerpb.RegisterResponse, error) {
	res := &registerpb.RegisterResponse{
		Response: fmt.Sprintf("Successfully registered"),
	}
	return res, nil
}
