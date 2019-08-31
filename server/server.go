package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/riyadennis/grpc-go-selfieday/api"
)

var collection *mongo.Collection

// Server is the grpc server and will hold data needed to
// perform the api calls.
type Server struct {
	dbCollection *mongo.Collection
}

// Settings holds the settings need to run the enpoints
type Settings struct {
	Server   *grpc.Server
	Listener net.Listener
	DBClient *mongo.Client
}

// NewSettings constructor for generating settings
func NewSettings() (*Settings, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:5052")
	if err != nil {
		return nil, err
	}
	mclient, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}
	return &Settings{
		Server:   grpc.NewServer(),
		Listener: lis,
		DBClient: mclient,
	}, nil
}

// Run initialises a grpc server
func (s *Settings) Run() {
	server := &Server{}
	api.RegisterRegisterServiceServer(s.Server, server)
	api.RegisterBlogServiceServer(s.Server, server)
	go func() {
		if err := s.Server.Serve(s.Listener); err != nil {
			panic(err)
		}
	}()

	mclient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mclient.Connect(ctx)
	server.dbCollection = mclient.Database("grpc_db").Collection("blog")

	s.cleanup(ctx)
}

// cleanup shuts down and closes all the resources
func (s *Settings) cleanup(ctx context.Context) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping server")
	s.Server.Stop()
	fmt.Println("closing the listener")
	s.Listener.Close()
	fmt.Println("closing mongo db connection")
	s.DBClient.Disconnect(ctx)
}
