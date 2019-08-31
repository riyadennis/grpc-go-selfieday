package main

import (
	"github.com/riyadennis/grpc-go-selfieday/server"
)

func main() {
	settings, err := server.NewSettings()
	if err != nil {
		panic(err)
	}
	settings.Run()
}
