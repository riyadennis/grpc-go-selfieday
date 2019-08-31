package server

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/riyadennis/grpc-go-selfieday/api"
)

// Blog holds the db structure for the blog
type Blog struct {
	ID       primitive.ObjectID `bson:"id, omitempty"`
	AuthorID string             `bson:"author_id, omitempty"`
	Title    string             `bson:"title,omitempty"`
	Content  string             `bson:"content,omitempty"`
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
