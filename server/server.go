package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/grpc-go-selfieday/api"
	"google.golang.org/grpc"
)

type server struct{}

// Registeration implement RegisterServiceServer
func (se *server) Registeration(ctx context.Context,
	req *api.RegisterRequest) (*api.RegisterResponse, error) {
	res := &api.RegisterResponse{
		Response: fmt.Sprintf("Registered %s", req.GetRegister().GetFirstName()),
	}
	return res, nil
}

// RegisterationManyTimes implement RegisterServiceServer for streaming
func (se *server) RegisterationManyTimes(req *api.RegisterStreamRequest,
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

func (se *server) ChatService(stream api.RegisterService_ChatServiceServer) error {
	fmt.Print("Chat service client streaming\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Print("finished streaming")
			response := &api.ChatResponse{
				Response: result,
			}
			return stream.SendAndClose(response)
		}
		if err != nil {
			panic(err)
		}
		message := req.GetChat().GetBody()
		result += message + " ! "
	}
	return nil
}

// Run initialises a grpc server
func Run() {
	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	api.RegisterRegisterServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
