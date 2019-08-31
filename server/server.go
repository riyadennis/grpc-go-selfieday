package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/riyadennis/grpc-go-selfieday/api"
)

var collection *mongo.Collection

// Server is the grpc server and will hold data needed to
// perform the api calls.
type Server struct {
	dbCollection *mongo.Collection
}

// Blog holds the db structure for the blog
type Blog struct {
	ID       primitive.ObjectID `bson:"id, omitempty"`
	AuthorID string             `bson:"author_id, omitempty"`
	Title    string             `bson:"title,omitempty"`
	Content  string             `bson:"content,omitempty"`
}

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

// CreateBlog will handle create blog request and will save the data to mongo db
func (se *Server) CreateBlog(ctx context.Context, req *api.CreateBlogRequest) (*api.CreateBlogResponse, error) {
	blog := req.Blog
	data := &Blog{
		AuthorID: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}
	res, err := se.dbCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error %v :: ", err))
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot find blog id "))
	}
	return &api.CreateBlogResponse{
		Blog: &api.Blog{
			Id:       id.Hex(),
			AuthorId: blog.AuthorId,
			Title:    blog.Title,
			Content:  blog.Content,
		},
	}, nil
}

// Run initialises a grpc server
func Run() {
	lis, err := net.Listen("tcp", "0.0.0.0:5052")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	server := &Server{}
	api.RegisterRegisterServiceServer(s, server)
	api.RegisterBlogServiceServer(s, server)
	go func() {
		if err := s.Serve(lis); err != nil {
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

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping server")
	s.Stop()
	fmt.Println("closing the listener")
	lis.Close()
	fmt.Println("closing mongo db connection")
	mclient.Disconnect(ctx)
}
