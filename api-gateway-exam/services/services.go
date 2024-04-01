package services

import (
	"exam/api-gateway/config"
	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"
	"exam/api-gateway/mock"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	UserMockService() mock.UserServiceClient
	CommentService() pbc.CommentServiceClient
	CommentMockService() mock.CommentServiceClient
	PostService() pbp.PostServiceClient
	PostMockService() mock.PostServiceClient
}

type serviceManager struct {
	userService        pbu.UserServiceClient
	userMockService    mock.UserServiceClient
	commentService     pbc.CommentServiceClient
	commentMockService mock.CommentServiceClient
	postService        pbp.PostServiceClient
	postMockService    mock.PostServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) UserMockService() mock.UserServiceClient {
	return s.userMockService
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func (s *serviceManager) PostMockService() mock.PostServiceClient {
	return s.postMockService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}

func (s *serviceManager) CommentMockService() mock.CommentServiceClient {
	return s.commentMockService
}

func NewServiceManager(cfg *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	//User service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	//Post service
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	//Comment service
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		// userMockService: mock.UserServiceClient{},
		// postMockService: mock.PostServiceClient{},
		// commentService: &mock.CommentServiceClient{},
		userService:    pbu.NewUserServiceClient(connUser),
		postService:    pbp.NewPostServiceClient(connPost),
		commentService: pbc.NewCommentServiceClient(connComment),
	}
	return serviceManager, nil
}
