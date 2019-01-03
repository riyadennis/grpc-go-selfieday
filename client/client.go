package main

import (
	"context"
	"fmt"
	"io"

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
