package mock

import (
	"context"
	pbp "exam/api-gateway/genproto/post_service"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type PostServiceClientI interface {
	Create(ctx context.Context, in *pbp.Post, opts ...grpc.CallOption) (*pbp.Post, error)
	Update(ctx context.Context, in *pbp.Post, opts ...grpc.CallOption) (*pbp.Post, error)
	Get(ctx context.Context, in *pbp.Id, opts ...grpc.CallOption) (*pbp.Post, error)
	Delete(ctx context.Context, in *pbp.Id, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *pbp.GetListFilter, opts ...grpc.CallOption) (*pbp.Posts, error)
}

type PostServiceClient struct {
}

func NewPostServiceClient() PostServiceClientI {
	return &PostServiceClient{}
}

func (c *PostServiceClient) Create(ctx context.Context, in *pbp.Post, opts ...grpc.CallOption) (*pbp.Post, error) {
	return in, nil
}

func (c *PostServiceClient) Update(ctx context.Context, in *pbp.Post, opts ...grpc.CallOption) (*pbp.Post, error) {
	return in, nil
}

func (c *PostServiceClient) Get(ctx context.Context, in *pbp.Id, opts ...grpc.CallOption) (*pbp.Post, error) {
	return &pbp.Post{
		Id:       "e292ca9d-d202-4aa2-a7de-487158b02dd4",
		UserId:   "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
		Content:  "Test Content 1",
		Likes:    10,
		Dislikes: 10,
		Views:    10,
		Category: "Test Category 1",
	}, nil
}

func (c *PostServiceClient) Delete(ctx context.Context, in *pbp.Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (c *PostServiceClient) List(ctx context.Context, in *pbp.GetListFilter, opts ...grpc.CallOption) (*pbp.Posts, error) {
	return &pbp.Posts{
		Items: []*pbp.Post{
			{
				Id:       "e292ca9d-d202-4aa2-a7de-487158b02dd1",
				UserId:   "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content:  "Test Content 1",
				Likes:    10,
				Dislikes: 10,
				Views:    10,
				Category: "Test Category 1",
			},
			{
				Id:       "e292ca9d-d202-4aa2-a7de-487158b02dd2",
				UserId:   "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content:  "Test Content 2",
				Likes:    10,
				Dislikes: 10,
				Views:    10,
				Category: "Test Category 2",
			},
			{
				Id:       "e292ca9d-d202-4aa2-a7de-487158b02dd3",
				UserId:   "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content:  "Test Content 3",
				Likes:    10,
				Dislikes: 10,
				Views:    10,
				Category: "Test Category 3",
			},
		},
		Count: 3,
	}, nil
}
