package repo

import (
	"context"
	pb "exam/user-service/genproto/user_service"

	"github.com/golang/protobuf/ptypes/empty"
)

// UserService interface
type UserServiceI interface {
	Create(ctx context.Context, req *pb.User) (*pb.User, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.UserModel, error)
	Update(ctx context.Context, req *pb.User) (*pb.User, error)
	Delete(ctx context.Context, req *pb.GetRequest) (*empty.Empty, error)
	List(ctx context.Context, req *pb.GetListFilter) (*pb.Users, error)
	CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.Status, error)
	UpdateRefresh(ctx context.Context, req *pb.UpdateRefreshReq) (*pb.User, error)
}
