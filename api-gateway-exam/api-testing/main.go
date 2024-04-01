package main

import (
	"exam/api-gateway/api-testing/handlers"
	"exam/api-gateway/api-testing/storage/kv"
	"log"
	"net/http"

	// "github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

func main() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// kv.Init(kv.NewRedisClient(client))
	kv.Init(kv.NewInMemoryInst())

	router := gin.New()

	router.POST("/users/register", handlers.RegisterUser)
	router.GET("/users/verify/:code", handlers.Verify)
	router.GET("/users/get", handlers.GetUser)
	router.POST("/users/create", handlers.CreateUser)
	router.DELETE("/users/delete", handlers.DeleteUser)
	router.GET("/users", handlers.ListUsers)

	router.GET("/posts/get", handlers.GetPost)
	router.POST("/posts/create", handlers.CreatePost)
	router.DELETE("/posts/delete", handlers.DeletePost)
	router.GET("/posts", handlers.ListPosts)

	router.GET("/comments/get", handlers.GetComment)
	router.POST("/comments/create", handlers.CreateComment)
	router.DELETE("/comments/delete", handlers.DeleteComment)
	router.GET("/comments", handlers.ListComments)

	log.Fatal(http.ListenAndServe(":9191", router))
}
