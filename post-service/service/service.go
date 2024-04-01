package service

import (
	"context"
	"database/sql"
	pb "template-post-service/genproto/post_service"
	"template-post-service/pkg/logger"
	"template-post-service/storage"

	// "github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostService
type PostService struct {
	storage storage.IStorage
	logger  logger.Logger
	// client  grpcclient.IServiceManager
}

// NewPostService
func NewPostService(collection *mongo.Collection, db *sql.DB, log logger.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(collection, db, log),
		logger:  log,
		// client:  client,
	}
}

// rpc Create(Post) returns (Post);
// rpc Update(Post) returns (Post);
// rpc Get(Post) returns (Post);
// rpc Delete(Id) returns (google.protobuf.Empty);
// rpc List(GetListFilter) returns (Posts);

func (p *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	return p.storage.Post().Create(req)
}

func (p *PostService) Get(ctx context.Context, id *pb.Id) (*pb.Post, error) {
	return p.storage.Post().Get(id)

}

func (p *PostService) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	return p.storage.Post().Update(req)
}

func (p *PostService) Delete(ctx context.Context, id *pb.Id) (*empty.Empty, error) {
	return p.storage.Post().Delete(id)
}

func (p *PostService) List(ctx context.Context, req *pb.GetListFilter) (*pb.Posts, error) {
	return p.storage.Post().List(req)
}
