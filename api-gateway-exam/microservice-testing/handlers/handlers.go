package handlers

import (
	"context"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"fmt"
	"net/http"

	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

type MockServiceManager struct {
	sm  services.IServiceManager
	log logger.Logger
}

func NewMockServiceManager(sm services.IServiceManager, log logger.Logger) *MockServiceManager {
	return &MockServiceManager{sm: sm, log: log}
}

func (sm *MockServiceManager) CreateUser(c *gin.Context) {
	var user pbu.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.UserService().Create(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while creating user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) GetUser(c *gin.Context) {
	id := c.Query("id")
	resp, err := sm.sm.UserService().Get(context.Background(), &pbu.GetRequest{UserId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while getting user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	_, err := sm.sm.UserService().Delete(context.Background(), &pbu.GetRequest{UserId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while deleting user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted",
	})
}

func (sm *MockServiceManager) UpdateUser(c *gin.Context) {
	var user pbu.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.UserService().Update(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while updating user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) ListUsers(c *gin.Context) {
	users, err := sm.sm.UserService().List(context.Background(), &pbu.GetListFilter{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while listing users", logger.Error(err))
		return
	}
	
	c.JSON(http.StatusOK, users)
}

// Post
func (sm *MockServiceManager) CreatePost(c *gin.Context) {
	var post pbp.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.PostService().Create(context.Background(), &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while creating post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) GetPost(c *gin.Context) {
	id := c.Query("id")

	resp, err := sm.sm.PostService().Get(context.Background(), &pbp.Id{PostId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while getting post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) DeletePost(c *gin.Context) {
	id := c.Query("id")

	_, err := sm.sm.PostService().Delete(context.Background(), &pbp.Id{PostId: id})
	fmt.Println(err, "error")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err, "---")

		sm.log.Error("error while deleting post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post was deleted",
	})
}

func (sm *MockServiceManager) UpdatePost(c *gin.Context) {
	var post pbp.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.PostService().Update(context.Background(), &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while updating post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) ListPosts(c *gin.Context) {
	posts, err := sm.sm.PostService().List(context.Background(), &pbp.GetListFilter{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while listing posts", logger.Error(err))
		return
	}
	
	c.JSON(http.StatusOK, posts)
}

// Comment
func (sm *MockServiceManager) CreateComment(c *gin.Context) {
	var comment pbc.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.CommentService().Create(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while creating comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) GetComment(c *gin.Context) {
	id := c.Query("id")

	resp, err := sm.sm.CommentService().Get(context.Background(), &pbc.Id{CommentId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while getting comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) DeleteComment(c *gin.Context) {
	id := c.Query("id")

	_, err := sm.sm.CommentService().Delete(context.Background(), &pbc.Id{CommentId: id})
	fmt.Println(err, " -- error here")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while deleting comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment was deleted",
	})
}

func (sm *MockServiceManager) UpdateComment(c *gin.Context) {
	var comment pbc.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.CommentService().Update(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while updating comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) ListComments(c *gin.Context) {
	comments, err := sm.sm.CommentService().List(context.Background(), &pbc.GetListFilter{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while listing comments", logger.Error(err))
		return
	}
	
	c.JSON(http.StatusOK, comments)
}