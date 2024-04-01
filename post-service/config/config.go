package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment       string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	MongoDBhost       string
	MongoDBport       string
	MongoDBdatabase   string
	MongoDBCollection string
	LogLevel          string
	RPCPort           string

	CassandraCluster      string
	CassandraKeyspaceName string

	// UserServiceHost string
	// UserServicePort int

	// CommentServiceHost string
	// CommentServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "socialdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "mubina2007"))

	//Mongo
	c.MongoDBhost = cast.ToString(getOrReturnDefault("MONGODB_HOST", "mongodb"))
	c.MongoDBport = cast.ToString(getOrReturnDefault("MONGODB_PORT", 27017))
	c.MongoDBdatabase = cast.ToString(getOrReturnDefault("MONGODB_DATABASE", "socialdb"))
	c.MongoDBCollection = cast.ToString(getOrReturnDefault("MONGODB_COLLECTION", "posts"))

	c.CassandraCluster = cast.ToString(getOrReturnDefault("CASSANDRA_CLUSTER", "cassandra"))
	c.CassandraKeyspaceName = cast.ToString(getOrReturnDefault("CASSANDRA_KEYSPACENAME", "my_application_keyspace"))

	// c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	// c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9090))

	// c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "localhost"))
	// c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 8080))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":7070"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return defaultValue
	}

	return defaultValue
}
