package main

import (
	"comment-service/config"
	pb "comment-service/genproto/comment_service"
	"comment-service/pkg/db"
	"comment-service/pkg/logger"

	// "comment-service/rabbitmq"
	"comment-service/service"
	"context"
	"fmt"
	"net"

	// "github.com/streadway/amqp"
	// "github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "comment-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, _, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("sql connection error", logger.Error(err))
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBhost, cfg.MongoDBport))
	clientMongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("failed to connect to mongodb: %v", logger.Error(err))
	}

	collection := clientMongo.Database(cfg.MongoDBdatabase).Collection(cfg.MongoDBCollection)
	// grpcCl, err := grpcClient.New(cfg)
	// if err != nil {
	// 	log.Fatal("cannot connect to grpc client:%v", logger.Error(err))
	// }

	// cluster := gocql.NewCluster(cfg.CassandraCluster)
	// cluster.Keyspace = cfg.CassandraKeyspaceName
	// cluster.Consistency = gocql.Quorum

	// // Create a session
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// // Create comments table
	// if err := session.Query(`
    //     CREATE TABLE IF NOT EXISTS comments (
    //         id UUID,
    //         post_id UUID,
    //         user_id UUID,
    //         content TEXT,
    //         created_at TIMESTAMP,
    //         updated_at TIMESTAMP,
    //         deleted_at TIMESTAMP,
    //         PRIMARY KEY (id, post_id, user_id)
    //     );
    // `).Exec(); err != nil {
	// 	panic(err)
	// }

	commentService := service.NewCommentService(collection, connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("cannot listen", logger.Error(err))
	}

	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// channel, err := conn.Channel()
	// if err != nil {
	// 	panic(err)
	// }

	// consumer := rabbitmq.NewRabbitMQConsumer(channel)
	// defer consumer.Close()

	// go func() {
	// 	consumer.ConsumeMessages(cfg.RabbitQueue, rabbitmq.ConsumeHandler, commentService)
	// }()

	server := grpc.NewServer()
	pb.RegisterCommentServiceServer(server, commentService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(lis); err != nil {
		log.Fatal("server cannot serve", logger.Error(err))
	}
}
