package server

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-go-selfieday/registerpb"
	"google.golang.org/grpc"
)

type server struct{}

// Registeration implement RegisterServiceServer
func (se *server) Registeration(ctx context.Context,
	req *registerpb.RegisterRequest) (*registerpb.RegisterResponse, error) {
	res := &registerpb.RegisterResponse{
		Response: fmt.Sprintf("Successfully registered"),
	}
	return res, nil
}

// RUN initialises a grpc server and add a listener to port
func Run() {
	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	registerpb.RegisterRegisterServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
