package main

import (
	"context"
	"fmt"

	"github.com/grpc-go-selfieday/registerpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := registerpb.NewRegisterServiceClient(conn)
	req := &registerpb.RegisterRequest{
		Register: &registerpb.Register{
			FirstName: "Peter",
			LastName:  "Donvon",
			Email:     "pete@down.com",
		},
	}
	resp, err := c.Registeration(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetResponse())
}
