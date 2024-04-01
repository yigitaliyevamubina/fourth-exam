package api

import (
	casb "exam/api-gateway/api/casbin"
	_ "exam/api-gateway/api/docs"
	v1 "exam/api-gateway/api/handlers/v1"
	"exam/api-gateway/api/handlers/v1/tokens"
	"exam/api-gateway/config"
	"exam/api-gateway/kafka/producer"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"exam/api-gateway/storage/mongorepo"
	"exam/api-gateway/storage/repo"

	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"

	admin "exam/api-gateway/storage/postgresrepo"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option Struct
type Option struct {
	InMemory       repo.InMemoryStorageI
	Cfg            config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Postgres       admin.AdminStorageI
	Producer       producer.KafkaProducer
	Mongo          mongorepo.AdminStorageI
}

// New -> constructor
// @title Social web
// @version 1.0
// @description Auth, Role-management, User, Post, Comment
// @host localhost:4040
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	casbinEnforcer, err := casbin.NewEnforcer(option.Cfg.AuthConfigPath, option.Cfg.AuthCSVPath)
	if err != nil {
		option.Logger.Error("cannot create a new enforcer", logger.Error(err))
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}

	casbinEnforcer.GetRoleManager().AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().AddMatchingFunc("keyMatch3", util.KeyMatch3)

	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")

	router.Use(cors.New(corsConfig))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandle := tokens.JWTHandler{
		SignInKey: option.Cfg.SignInKey,
		Log:       option.Logger,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		InMemoryStorage: option.InMemory,
		Log:             option.Logger,
		ServiceManager:  option.ServiceManager,
		Cfg:             option.Cfg,
		JWTHandler:      jwtHandle,
		Postgres:        option.Postgres,
		Casbin:          casbinEnforcer,
		Producer:        option.Producer,
		Mongo:           option.Mongo,
	})

	api := router.Group("/v1")

	api.Use(casb.NewAuth(casbinEnforcer, option.Cfg))

	//insertion
	api.GET("/insert", handlerV1.InsertToMongo) //unauthorized

	//rbac
	api.GET("/rbac/roles", handlerV1.ListAllRoles)              //superadmin
	api.GET("/rbac/policies/:role", handlerV1.ListRolePolicies) //superadmin
	api.POST("/rbac/add/policy", handlerV1.AddPolicyToRole)     //superadmin
	api.DELETE("/rbac/delete/policy", handlerV1.DeletePolicy)   //superadmin

	//users
	api.POST("/user/create", handlerV1.CreateUser)          //admin
	api.POST("/user/register", handlerV1.Register)          //unauthorized
	api.GET("/user", handlerV1.GetUserById)                 //user
	api.PUT("/user/update/:id", handlerV1.UpdateUser)       //user
	api.DELETE("/user/delete", handlerV1.DeleteUser)        //user
	api.GET("/users/:page/:limit", handlerV1.GetAllUsers)   //user
	api.GET("/user/verify/:email/:code", handlerV1.Verify)  //unauthorized
	api.POST("/user/login", handlerV1.Login)                //unauthorized
	api.POST("/user/refresh", handlerV1.UpdateRefreshToken) //user

	//post
	api.POST("/post/create", handlerV1.CreatePost)                       //user
	api.PUT("/post/update/:id", handlerV1.UpdatePost)                    //user
	api.GET("/post/:id", handlerV1.GetPostById)                          //user
	api.DELETE("/post/delete/:id", handlerV1.DeletePost)                 //user
	api.GET("/posts/:page/:limit", handlerV1.ListPosts)                  //user
	api.GET("/posts/:page/:limit/:user_id", handlerV1.ListPostsByUserId) //user
	api.PUT("/post/like", handlerV1.LikePost)                            //user
	api.PUT("/post/dislike", handlerV1.DislikePost)                      //user

	//comment
	api.POST("/comment/create", handlerV1.CreateComment)                           //user
	api.PUT("/comment/update/:id", handlerV1.UpdateComment)                        //user
	api.GET("/comment/:id", handlerV1.GetCommentById)                              //user
	api.DELETE("/comment/delete/:id", handlerV1.DeleteComment)                     //user
	api.GET("/comments/:page/:limit", handlerV1.ListComments)                      //user
	api.GET("/comments/:page/:limit/:user_id", handlerV1.ListCommentsByUserId)     //user
	api.GET("/get/comments/:page/:limit/:post_id", handlerV1.ListCommentsByPostId) //user

	//admin
	api.POST("/auth/create", handlerV1.CreateAdmin)   //superadmin
	api.DELETE("/auth/delete", handlerV1.DeleteAdmin) //superadmin
	api.POST("/auth/login", handlerV1.LoginAdmin)     //unauthorized

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
