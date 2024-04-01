package storage

import (
	"exam/user-service/pkg/db"
	"exam/user-service/pkg/logger"

	"exam/user-service/storage/postgres"

	// postgres "exam/user-service/storage/cassandra"
	// "exam/user-service/storage/mongodb"
	"exam/user-service/storage/repo"

	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
)

// Storage
type StorageI interface {
	UserService() repo.UserServiceI
}

type storagePg struct {
	userService repo.UserServiceI
}

func New(collection *mongo.Collection, db *db.Postgres, log logger.Logger) StorageI {
	return &storagePg{userService: postgres.NewUserRepo(db, log)}
	// return &storagePg{userService: mongodb.NewUserRepo(collection, log)}
	// return &storagePg{userService: postgres.NewCassandraRepo(&cluster)}

}

func (s *storagePg) UserService() repo.UserServiceI {
	return s.userService
}
