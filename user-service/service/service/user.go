package service

import (
	"context"
	pbc "exam/user-service/genproto/comment_service"
	pbp "exam/user-service/genproto/post_service"
	pb "exam/user-service/genproto/user_service"
	"exam/user-service/pkg/logger"
	grpcClient "exam/user-service/service/grpc_client"
	"exam/user-service/storage"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	storage storage.StorageI
	log     logger.Logger
	service grpcClient.IServiceManager
}

// Constructor
func NewUserService(storage storage.StorageI, log logger.Logger, service grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage,
		log:     log,
		service: service,
	}
}

func (c *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return c.storage.UserService().Create(ctx, req)
}

func (c *UserService) Get(ctx context.Context, req *pb.GetRequest) (*pb.UserModel, error) { 
	user, err := c.storage.UserService().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	posts, err := c.service.PostService().List(ctx, &pbp.GetListFilter{Page: 1, Limit: 100000, UserId: user.Id})
	if err != nil {
		c.log.Error("error while getting posts by user id: %v", logger.Error(err))
		return nil, err
	}

	userPosts := make([]*pb.Post, 0, len(posts.Items))

	for _, post := range posts.Items {
		comments, err := c.service.CommentService().List(ctx, &pbc.GetListFilter{Page: 1, Limit: 10000, PostId: post.Id})
		if err != nil {
			c.log.Error("error while getting comments by post id: %v", logger.Error(err))
		}

		postComments := make([]*pb.Comment, 0, len(comments.Items))

		for _, comment := range comments.Items {
			owner, err := c.storage.UserService().Get(ctx, &pb.GetRequest{UserId: comment.UserId})
			if err != nil {
				c.log.Error("error while getting comment's owner by user id: %v", logger.Error(err))
			}

			userComment := &pb.Comment{
				Id:      comment.Id,
				PostId:  comment.PostId,
				UserId:  comment.UserId,
				Content: comment.Content,
				Owner: &pb.User{
					Id:           owner.Id,
					Username:     owner.Username,
					FirstName:    owner.FirstName,
					LastName:     owner.LastName,
					Email:        owner.Email,
					Password:     owner.Password,
					Bio:          owner.Bio,
					Website:      owner.Website,
					CreatedAt:    owner.CreatedAt,
					UpdatedAt:    owner.UpdatedAt,
					RefreshToken: owner.RefreshToken,
				},
			}
			postComments = append(postComments, userComment)
		}

		userPost := &pb.Post{
			Id:       post.Id,
			UserId:   post.UserId,
			Content:  post.Content,
			Likes:    post.Likes,
			Dislikes: post.Dislikes,
			Views:    post.Views,
			Category: post.Category,
			Comments: postComments,
		}

		userPosts = append(userPosts, userPost)
	}

	user.Posts = userPosts

	return user, nil
}

func (c *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	return c.storage.UserService().Update(ctx, req)
}

func (c *UserService) Delete(ctx context.Context, req *pb.GetRequest) (*empty.Empty, error) {
	return c.storage.UserService().Delete(ctx, req)
}

func (c *UserService) List(ctx context.Context, req *pb.GetListFilter) (*pb.Users, error) {
	users, err := c.storage.UserService().List(ctx, req)
	if err != nil {
		return nil, err
	}

	usersResp := make([]*pb.UserModel, 0, len(users.Users))

	for _, user := range users.Users {
		posts, err := c.service.PostService().List(ctx, &pbp.GetListFilter{Page: 1, Limit: 100000, UserId: user.Id})
		if err != nil {
			c.log.Error("error while getting posts by user id: %v", logger.Error(err))
			return nil, err
		}

		userPosts := make([]*pb.Post, 0, len(posts.Items))

		for _, post := range posts.Items {
			comments, err := c.service.CommentService().List(ctx, &pbc.GetListFilter{Page: 1, Limit: 10000, PostId: post.Id})
			if err != nil {
				c.log.Error("error while getting comments by post id: %v", logger.Error(err))
			}

			postComments := make([]*pb.Comment, 0, len(comments.Items))

			for _, comment := range comments.Items {
				owner, err := c.storage.UserService().Get(ctx, &pb.GetRequest{UserId: comment.UserId})
				if err != nil {
					c.log.Error("error while getting comment's owner by user id: %v", logger.Error(err))
				}

				userComment := &pb.Comment{
					Id:      comment.Id,
					PostId:  comment.PostId,
					UserId:  comment.UserId,
					Content: comment.Content,
					Owner: &pb.User{
						Id:           owner.Id,
						Username:     owner.Username,
						FirstName:    owner.FirstName,
						LastName:     owner.LastName,
						Email:        owner.Email,
						Password:     owner.Password,
						Bio:          owner.Bio,
						Website:      owner.Website,
						CreatedAt:    owner.CreatedAt,
						UpdatedAt:    owner.UpdatedAt,
						RefreshToken: owner.RefreshToken,
					},
				}
				postComments = append(postComments, userComment)
			}

			userPost := &pb.Post{
				Id:       post.Id,
				UserId:   post.UserId,
				Content:  post.Content,
				Likes:    post.Likes,
				Dislikes: post.Dislikes,
				Views:    post.Views,
				Category: post.Category,
				Comments: postComments,
			}

			userPosts = append(userPosts, userPost)
		}

		user.Posts = userPosts
		usersResp = append(usersResp, user)
	}

	users.Users = usersResp
	return users, nil
}

func (c *UserService) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.Status, error) {
	return c.storage.UserService().CheckField(ctx, req)
}

func (c *UserService) UpdateRefresh(ctx context.Context, req *pb.UpdateRefreshReq) (*pb.User, error) {
	return c.storage.UserService().UpdateRefresh(ctx, req)
}
