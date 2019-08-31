package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/riyadennis/grpc-go-selfieday/api"
)

// Registeration implement RegisterServiceServer
func (se *Server) Registeration(ctx context.Context,
	req *api.RegisterRequest) (*api.RegisterResponse, error) {
	res := &api.RegisterResponse{
		Response: fmt.Sprintf("Registered %s", req.GetRegister().GetFirstName()),
	}
	return res, nil
}

// RegisterationManyTimes implement RegisterServiceServer for streaming
func (se *Server) RegisterationManyTimes(req *api.RegisterStreamRequest,
	stream api.RegisterService_RegisterationManyTimesServer) error {
	fname := req.GetRegister().GetFirstName()
	for i := 0; i < 10; i++ {
		res := &api.RegisterStreamResponse{
			Response: fmt.Sprintf("Hello %s number %s", fname, strconv.Itoa(i)),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
