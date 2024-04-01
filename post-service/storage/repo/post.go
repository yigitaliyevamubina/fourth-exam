package repo

import (
	pb "template-post-service/genproto/post_service"

	"github.com/golang/protobuf/ptypes/empty"
)

// rpc Create(Post) returns (Post);
//	rpc Update(Post) returns (Post);
//	rpc Get(Post) returns (Post);
//	rpc Delete(Id) returns (google.protobuf.Empty);
//	rpc List(GetListFilter) returns (Posts);

// PostStorageI
type PostStorageI interface {
	Create(*pb.Post) (*pb.Post, error)
	Get(*pb.Id) (*pb.Post, error)
	Update(*pb.Post) (*pb.Post, error)
	Delete(*pb.Id) (*empty.Empty, error)
	List(*pb.GetListFilter) (*pb.Posts, error)
}
