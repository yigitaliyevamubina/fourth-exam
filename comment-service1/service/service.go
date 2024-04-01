package service

import (
	pb "comment-service/genproto/comment_service"
	"comment-service/pkg/logger"
	"comment-service/storage"
	"context"
	"database/sql"

	// "github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentService struct {
	storage storage.IStorage
	logger  logger.Logger
	// service grpcClient.IServiceManager
}

func NewCommentService(collection *mongo.Collection, db *sql.DB, log logger.Logger) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(collection, db, log),
		logger:  log,
		// service: service,
	}
}

// Create(req *pb.Comment) (*pb.Comment, error)
// Update(req *pb.Comment) (*pb.Comment, error)
// Get(id *pb.Id) (*pb.Comment, error)
// Delete(id *pb.Id) (*empty.Empty, error)
// List(Greq *pb.GetListFilter) (*pb.Comments, error)

func (c *CommentService) Create(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	return c.storage.Comment().Create(req)
}

func (c *CommentService) Get(ctx context.Context, id *pb.Id) (*pb.Comment, error) {
	return c.storage.Comment().Get(id)
}

func (c *CommentService) Update(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	return c.storage.Comment().Update(req)
}

func (c *CommentService) Delete(ctx context.Context, id *pb.Id) (*empty.Empty, error) {
	return c.storage.Comment().Delete(id)
}

func (c *CommentService) List(ctx context.Context, req *pb.GetListFilter) (*pb.Comments, error) {
	return c.storage.Comment().List(req)

}
