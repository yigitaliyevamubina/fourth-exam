package storage

import (
	"database/sql"
	"template-post-service/pkg/logger"

	"template-post-service/storage/postgres"
	// "template-post-service/storage/cassandra"
	// "template-post-service/storage/mongodb"
	"template-post-service/storage/repo"

	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sql.DB
	postRepo repo.PostStorageI
}

func NewStoragePg(collection *mongo.Collection, db *sql.DB, log logger.Logger) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
		// postRepo: mongodb.NewPostRepo(collection, log),
		// postRepo: cassandra.NewPostRepo(cluster),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
