package main

import (
	"context"
	"exam/api-gateway/api"
	"exam/api-gateway/config"
	"exam/api-gateway/kafka/producer"
	"exam/api-gateway/pkg/db"
	"exam/api-gateway/pkg/etc"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"exam/api-gateway/storage/mongodb"
	"exam/api-gateway/storage/postgres"
	"exam/api-gateway/storage/redis"
	"fmt"

	rds "github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//login superadmin -> username = 'a' -> password = 'b'

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	fmt.Println(etc.HashPassword("b"))

	redisPool := rds.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (rds.Conn, error) {
			c, err := rds.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	db, _, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("cannot run http server", logger.Error(err))
		panic(err)
	}

	writer, err := producer.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("cannot create a kafka producer", logger.Error(err))
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBhost, cfg.MongoDBport))
	clientMongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("failed to connect to mongodb: %v", logger.Error(err))
	}

	collection := clientMongo.Database(cfg.MongoDBdatabase).Collection(cfg.MongoDBCollection)

	server := api.New(api.Option{
		InMemory:       redis.NewRedisRepo(&redisPool),
		Cfg:            cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		Postgres:       postgres.NewAdminRepo(db),
		Producer:       writer,
		Mongo:          mongodb.NewAdminRepo(collection),
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("cannot run http server", logger.Error(err))
		panic(err)
	}
}
