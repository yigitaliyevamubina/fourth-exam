package grpcClient

import (
	"comment-service/config"
)

type IServiceManager interface {
	// PostService() pbp.PostServiceClient
	// UserService() pbu.UserServiceClient
}

type serviceManager struct {
	cfg config.Config
	// postService pbp.PostServiceClient
	// userService pbu.UserServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// connPost, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("error while dialing to the post service", logger.Error(err))
	// }

	// connUser, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("error while dialing to the user service", logger.Error(err))
	// }

	return &serviceManager{
		cfg: cfg,
		// postService: pbp.NewPostServiceClient(connPost),
		// userService: pbu.NewUserServiceClient(connUser),
	}, nil
}

// func (s *serviceManager) PostService() pbp.PostServiceClient {
// 	return s.postService
// }

// func (s *serviceManager) UserService() pbu.UserServiceClient {
// 	return s.userService
// }
