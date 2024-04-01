package storage

import (
	"comment-service/pkg/logger"
	"comment-service/storage/postgres"
	"comment-service/storage/repo"
	"database/sql"

	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db          *sql.DB
	commentRepo repo.CommentStorageI
}

func NewStoragePg(collection *mongo.Collection, db *sql.DB, log logger.Logger) *storagePg {
	return &storagePg{
		db:          db,
		commentRepo: postgres.NewCommentRepo(db),
		// commentRepo: mongodb.NewCommentRepo(collection, log),
		// commentRepo: cassandra.NewCommentRepo(cluster),
	}
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}
