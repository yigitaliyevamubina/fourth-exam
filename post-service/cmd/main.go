package main

import (
	"context"
	"fmt"
	"net"
	config "template-post-service/config"
	pb "template-post-service/genproto/post_service"
	"template-post-service/pkg/db"
	"template-post-service/pkg/logger"

	"template-post-service/service"

	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, _, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres error", logger.Error(err))
	}

	// client, err := grpcCLient.New(cfg)
	// if err != nil {
	// 	log.Fatal("error while adding grpc client", logger.Error(err))
	// }

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBhost, cfg.MongoDBport))
	clientMongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("failed to connect to mongodb: %v", logger.Error(err))
	}

	collection := clientMongo.Database(cfg.MongoDBdatabase).Collection(cfg.MongoDBCollection)

	// cluster := gocql.NewCluster(cfg.CassandraCluster)
	// cluster.Keyspace = cfg.CassandraKeyspaceName
	// cluster.Consistency = gocql.Quorum

	// // Create a session
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// // Create posts table
	// if err := session.Query(`
    //     CREATE TABLE IF NOT EXISTS posts (
    //         id UUID PRIMARY KEY,
    //         user_id UUID,
    //         content TEXT,
    //         title TEXT,
    //         likes BIGINT,
    //         dislikes BIGINT,
    //         views BIGINT,
    //         category TEXT,
    //         created_at TIMESTAMP,
    //         updated_at TIMESTAMP,
    //         deleted_at TIMESTAMP
    //     );
    // `).Exec(); err != nil {
	// 	panic(err)
	// }

	postService := service.NewPostService(collection, connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("failed to listen to: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server is running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}
}
