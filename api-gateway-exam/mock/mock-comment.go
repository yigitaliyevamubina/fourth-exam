package mock

import (
	"context"
	pbc "exam/api-gateway/genproto/comment_service"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type CommentServiceClientI interface {
	Create(ctx context.Context, in *pbc.Comment, opts ...grpc.CallOption) (*pbc.Comment, error)
	Update(ctx context.Context, in *pbc.Comment, opts ...grpc.CallOption) (*pbc.Comment, error)
	Get(ctx context.Context, in *pbc.Id, opts ...grpc.CallOption) (*pbc.Comment, error)
	Delete(ctx context.Context, in *pbc.Id, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *pbc.GetListFilter, opts ...grpc.CallOption) (*pbc.Comments, error)
}

type CommentServiceClient struct {
}

func NewCommentServiceClient() CommentServiceClientI {
	return &CommentServiceClient{}
}

func (c *CommentServiceClient) Create(ctx context.Context, in *pbc.Comment, opts ...grpc.CallOption) (*pbc.Comment, error) {
	return in, nil
}

func (c *CommentServiceClient) Update(ctx context.Context, in *pbc.Comment, opts ...grpc.CallOption) (*pbc.Comment, error) {
	return in, nil
}

func (c *CommentServiceClient) Get(ctx context.Context, in *pbc.Id, opts ...grpc.CallOption) (*pbc.Comment, error) {
	return &pbc.Comment{
		Id:      "",
		PostId:  "e292ca9d-d202-4aa2-a7de-487158b02dd4",
		UserId:  "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
		Content: "Test Content 1",
	}, nil
}

func (c *CommentServiceClient) Delete(ctx context.Context, in *pbc.Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (c *CommentServiceClient) List(ctx context.Context, in *pbc.GetListFilter, opts ...grpc.CallOption) (*pbc.Comments, error) {
	return &pbc.Comments{
		Count: 3,
		Items: []*pbc.Comment{
			{
				Id:      "",
				PostId:  "e292ca9d-d202-4aa2-a7de-487158b02dd4",
				UserId:  "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content: "Test Content 1",
			},
			{
				Id:      "",
				PostId:  "e292ca9d-d202-4aa2-a7de-487158b02dd4",
				UserId:  "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content: "Test Content 2",
			},
			{
				Id:      "",
				PostId:  "e292ca9d-d202-4aa2-a7de-487158b02dd4",
				UserId:  "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				Content: "Test Content 3",
			},
		},
	}, nil
}
