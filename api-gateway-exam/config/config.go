package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string //develop, staging, production

	RedisHost string
	RedisPort int

	UserServiceHost string
	UserServicePort int

	PostServiceHost string
	PostServicePort int

	CommentServiceHost string
	CommentServicePort int

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	//context timeout in seconds
	CtxTimeOut int

	LogLevel string
	HTTPPort string

	AccessTokenTimeout  int //minutes
	RefreshTokenTimeout int //hours
	AuthConfigPath      string
	AuthCSVPath         string

	MongoDBhost       string
	MongoDBport       string
	MongoDBdatabase   string
	MongoDBCollection string

	SignInKey string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":4040"))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "redisdb"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "socialdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "mubina2007"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "userservice"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9090))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "postservice"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 7070))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "commentservice"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 8080))

	c.AccessTokenTimeout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMEOUT", 1000))
	c.RefreshTokenTimeout = cast.ToInt(getOrReturnDefault("REFRESH_TOKEN_TIMEOUT", 1000))

	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATH", "./config/auth.conf"))
	c.AuthCSVPath = cast.ToString(getOrReturnDefault("AUTH_CSV_PATH", "./config/auth.csv"))
	//Mongo
	c.MongoDBhost = cast.ToString(getOrReturnDefault("MONGODB_HOST", "mongodb"))
	c.MongoDBport = cast.ToString(getOrReturnDefault("MONGODB_PORT", 27017))
	c.MongoDBdatabase = cast.ToString(getOrReturnDefault("MONGODB_DATABASE", "socialdb"))
	c.MongoDBCollection = cast.ToString(getOrReturnDefault("MONGODB_COLLECTION", "admins"))

	c.CtxTimeOut = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	c.SignInKey = cast.ToString(getOrReturnDefault("SIGN_IN_KEY", "golang-exam-4d-month"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return defaultValue
	}

	return defaultValue
}
