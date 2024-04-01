package mock

import (
	"context"
	pbu "exam/api-gateway/genproto/user_service"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type UserServiceClientI interface {
	Create(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	Get(ctx context.Context, in *pbu.GetRequest, opts ...grpc.CallOption) (*pbu.User, error)
	Update(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	Delete(ctx context.Context, in *pbu.GetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *pbu.GetListFilter, opts ...grpc.CallOption) (*pbu.Users, error)
	CheckField(ctx context.Context, in *pbu.CheckFieldReq, opts ...grpc.CallOption) (*pbu.Status, error)
	UpdateRefresh(ctx context.Context, in *pbu.UpdateRefreshReq, opts ...grpc.CallOption) (*pbu.User, error)
}

type UserServiceClient struct {
}

func NewUserServiceClient() UserServiceClientI {
	return &UserServiceClient{}
}

func (c *UserServiceClient) Create(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *UserServiceClient) Get(ctx context.Context, in *pbu.GetRequest, opts ...grpc.CallOption) (*pbu.User, error) {
	return &pbu.User{
		Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb1aaf",
		FirstName:    "Test FirstName 1",
		LastName:     "Test Lastname 1",
		Username:     "Test Username 1",
		Bio:          "Test bio 1",
		Website:      "Test website 1",
		Email:        "testemail@gmail.com",
		Password:     "**kkw##knrtest",
		RefreshToken: "refresh token test",
		CreatedAt:    time.Now().String(),
	}, nil
}

func (c *UserServiceClient) Update(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *UserServiceClient) Delete(ctx context.Context, in *pbu.GetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (c *UserServiceClient) List(ctx context.Context, in *pbu.GetListFilter, opts ...grpc.CallOption) (*pbu.Users, error) {
	return &pbu.Users{
		Count: 3,
		Users: []*pbu.UserModel{
			{
				Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb1aaf",
				FirstName:    "Test FirstName 1",
				LastName:     "Test Lastname 1",
				Username:     "Test Username 1",
				Bio:          "Test bio 1",
				Website:      "Test website 1",
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest",
				RefreshToken: "refresh token test",
				CreatedAt:    time.Now().String(),
			},
			{
				Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				FirstName:    "Test FirstName 2",
				LastName:     "Test Lastname 2",
				Username:     "Test Username 2",
				Bio:          "Test bio 2",
				Website:      "Test website 2",
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest",
				RefreshToken: "refresh token test",
				CreatedAt:    time.Now().String(),
			},
			{
				Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb3aaf",
				FirstName:    "Test FirstName 3",
				LastName:     "Test Lastname 3",
				Username:     "Test Username 3",
				Bio:          "Test bio 3",
				Website:      "Test website 3",
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest",
				RefreshToken: "refresh token test",
				CreatedAt:    time.Now().String(),
			},
		},
	}, nil
}

func (c *UserServiceClient) CheckField(ctx context.Context, in *pbu.CheckFieldReq, opts ...grpc.CallOption) (*pbu.Status, error) {
	return &pbu.Status{Status: true}, nil
}

func (c *UserServiceClient) UpdateRefresh(ctx context.Context, in *pbu.UpdateRefreshReq, opts ...grpc.CallOption) (*pbu.User, error) {
	return &pbu.User{
		Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb1aaf",
		FirstName:    "Test FirstName 1",
		LastName:     "Test Lastname 1",
		Username:     "Test Username 1",
		Bio:          "Test bio 1",
		Website:      "Test website 1",
		Email:        "testemail@gmail.com",
		Password:     "**kkw##knrtest",
		RefreshToken: "refresh token test",
		CreatedAt:    time.Now().String(),
	}, nil
}
