package repo

import (
	pb "comment-service/genproto/comment_service"

	"github.com/golang/protobuf/ptypes/empty"
)

type CommentStorageI interface {
	Create(req *pb.Comment) (*pb.Comment, error)
	Update(req *pb.Comment) (*pb.Comment, error)
	Get(id *pb.Id) (*pb.Comment, error)
	Delete(id *pb.Id) (*empty.Empty, error)
	List(Greq *pb.GetListFilter) (*pb.Comments, error)
}
