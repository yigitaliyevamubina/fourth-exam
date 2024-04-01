package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment        string // develop, staging, production
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	RPCPort            string
	PostServiceHost    string
	PostServicePort    int
	CommentServiceHost string
	CommentServicePort int
	MongoDBhost        string
	MongoDBport        string
	MongoDBdatabase    string
	MongoDBCollection  string

	CassandraCluster      string
	CassandraKeyspaceName string

}

func Load() *Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "socialdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "mubina2007"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9090"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "postservice"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", "7070"))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "commentservice"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", "8080"))

	//Mongo
	c.MongoDBhost = cast.ToString(getOrReturnDefault("MONGODB_HOST", "mongodb"))
	c.MongoDBport = cast.ToString(getOrReturnDefault("MONGODB_PORT", 27017))
	c.MongoDBdatabase = cast.ToString(getOrReturnDefault("MONGODB_DATABASE", "socialdb"))
	c.MongoDBCollection = cast.ToString(getOrReturnDefault("MONGODB_COLLECTION", "users"))

	c.CassandraCluster = cast.ToString(getOrReturnDefault("CASSANDRA_CLUSTER", "cassandra"))
	c.CassandraKeyspaceName = cast.ToString(getOrReturnDefault("CASSANDRA_KEYSPACENAME", "my_application_keyspace"))

	return &c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return defaultValue
	}

	return defaultValue
}
