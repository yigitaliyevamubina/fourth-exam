package grpc_client

import (
	"template-post-service/config"
)

type IServiceManager interface {
	// UserService() pbu.UserServiceClient
	// CommentService() pbc.CommentServiceClient
}

type serviceManager struct {
	cfg config.Config
	// userService    pbu.UserServiceClient
	// commentService pbc.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// connComment, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("error while dialing to the comment service", logger.Error(err))
	// }

	// connUser, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("error while dialing to the user service", logger.Error(err))
	// }

	return &serviceManager{
		cfg: cfg,
		// commentService: pbc.NewCommentServiceClient(connComment),
		// userService:    pbu.NewUserServiceClient(connUser),
	}, nil
}

// func (s *serviceManager) CommentService() pbc.CommentServiceClient {
// 	return s.commentService
// }

// func (s *serviceManager) UserService() pbu.UserServiceClient {
// 	return s.userService
// }
