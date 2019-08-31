package server

import (
	"fmt"
	"io"

	"github.com/riyadennis/grpc-go-selfieday/api"
)

// ChatService creates a stream of chat messages
func (se *Server) ChatService(stream api.RegisterService_ChatServiceServer) error {
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
}
