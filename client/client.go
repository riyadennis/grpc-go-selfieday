package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/grpc-go-selfieday/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := api.NewRegisterServiceClient(conn)
	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
}

// doUnary send one request to the server
func doUnary(cc api.RegisterServiceClient) {
	req := &api.RegisterRequest{
		Register: &api.Register{
			FirstName: "Peter",
			LastName:  "Donvon",
			Email:     "pete@down.com",
		},
	}
	resp, err := cc.Registeration(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetResponse())
}

// doServerStreaming send one request to the server
func doServerStreaming(cc api.RegisterServiceClient) {
	req := &api.RegisterStreamRequest{
		Register: &api.Register{
			FirstName: "Peter",
			LastName:  "Donvon",
			Email:     "pete@down.com",
		},
	}
	stream, err := cc.RegisterationManyTimes(context.Background(), req)
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("response %v \n", res.GetResponse())
	}
}
func doClientStreaming(cc api.RegisterServiceClient) {
	fmt.Print("Doing client streaming")
	stream, err := cc.ChatService(context.Background())
	if err != nil {
		panic(err)
	}
	req := []*api.ChatRequest{
		&api.ChatRequest{
			Chat: &api.Chat{
				Id:      1,
				Heading: "first message",
				Body:    "Hello from the other side",
			},
		},
		&api.ChatRequest{
			Chat: &api.Chat{
				Id:      2,
				Heading: "second message",
				Body:    " second Hello from the other side",
			},
		},
		&api.ChatRequest{
			Chat: &api.Chat{
				Id:      3,
				Heading: "third message",
				Body:    " third Hello from the other side",
			},
		},
		&api.ChatRequest{
			Chat: &api.Chat{
				Id:      4,
				Heading: "fourth message",
				Body:    "fourth Hello from the other side",
			},
		},
	}
	for _, r := range req {
		stream.Send(r)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response %v", res)
}
