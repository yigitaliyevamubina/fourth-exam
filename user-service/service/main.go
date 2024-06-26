package service

import (
	"context"
	"exam/user-service/config"
	pb "exam/user-service/genproto/user_service"
	"exam/user-service/pkg/db"
	"exam/user-service/pkg/logger"
	grpcClient2 "exam/user-service/service/grpc_client"
	"exam/user-service/service/service"
	storage2 "exam/user-service/storage"
	"fmt"
	"net"

	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Service struct {
	UserService *service.UserService
}

func New(cfg *config.Config, log logger.Logger) (*Service, error) {
	postgres, err := db.New(*cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database:%v", err.Error())
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBhost, cfg.MongoDBport))
	clientMongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("failed to connect to mongodb: %v", logger.Error(err))
	}

	// cluster := gocql.NewCluster(cfg.CassandraCluster)
	// cluster.Keyspace = cfg.CassandraKeyspaceName
	// cluster.Consistency = gocql.Quorum

	// // Create a session
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// // defer session.Close()

	// // Create users table
	// if err := session.Query(`
    //     CREATE TABLE IF NOT EXISTS users (
    //         id UUID PRIMARY KEY,
    //         username TEXT,
    //         email TEXT,
    //         password TEXT,
    //         first_name TEXT,
    //         last_name TEXT,
    //         bio TEXT,
    //         website TEXT,
    //         is_active BOOLEAN,
    //         refresh_token TEXT,
    //         created_at TIMESTAMP,
    //         updated_at TIMESTAMP,
    //         deleted_at TIMESTAMP
    //     );
    // `).Exec(); err != nil {
	// 	panic(err)
	// }

	collection := clientMongo.Database(cfg.MongoDBdatabase).Collection(cfg.MongoDBCollection)
	storage := storage2.New(collection, postgres, log)
	grpcClient, err := grpcClient2.New(*cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to grpc client:%v", err.Error())
	}

	return &Service{UserService: service.NewUserService(storage, log, grpcClient)}, nil
}

func (s *Service) Run(log logger.Logger, cfg *config.Config) {
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, s.UserService)

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while creating a listener", logger.Error(err))
		return
	}

	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
		logger.String("rpc port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
