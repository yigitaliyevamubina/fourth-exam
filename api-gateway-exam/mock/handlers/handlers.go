package handlers

import (
	"context"
	"exam/api-gateway/mock"
	"net/http"

	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService    *mock.UserServiceClient
	PostService    *mock.PostServiceClient
	CommentService *mock.CommentServiceClient
}

func NewHandler(userService *mock.UserServiceClient, postService *mock.PostServiceClient, commentService *mock.CommentServiceClient) *Handler {
	return &Handler{UserService: userService, PostService: postService, CommentService: commentService}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser pbu.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.Create(context.Background(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var newUser pbu.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.Update(context.Background(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Query("id")

	user, err := h.UserService.Get(context.Background(), &pbu.GetRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Query("id")

	_, err := h.UserService.Delete(context.Background(), &pbu.GetRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.UserService.List(context.Background(), &pbu.GetListFilter{Page: 1, Limit: 10, OrderBy: "created_at"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) CheckField(c *gin.Context) {
	var body pbu.CheckFieldReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, err := h.UserService.CheckField(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if status.Status {
		c.JSON(http.StatusOK, gin.H{
			"message": "user exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user does not exist",
	})
}

func (h *Handler) UpdateRefresh(c *gin.Context) {
	var body pbu.UpdateRefreshReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.UpdateRefresh(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Post handlers
func (h *Handler) CreatePost(c *gin.Context) {
	var post pbp.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postResp, err := h.PostService.Create(context.Background(), &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, postResp)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	var post pbp.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postRes, err := h.PostService.Update(context.Background(), &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, postRes)
}

func (h *Handler) GetPost(c *gin.Context) {
	postId := c.Query("id")

	product, err := h.PostService.Get(context.Background(), &pbp.Id{PostId: postId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) DeletePost(c *gin.Context) {
	postId := c.Query("id")

	_, err := h.PostService.Delete(context.Background(), &pbp.Id{PostId: postId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post was deleted successfully",
	})
}

func (h *Handler) ListPosts(c *gin.Context) {
	users, err := h.PostService.List(context.Background(), &pbp.GetListFilter{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Comment handlers
func (h *Handler) CreateComment(c *gin.Context) {
	var comment pbc.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentResp, err := h.CommentService.Create(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commentResp)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var comment pbc.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentRes, err := h.CommentService.Update(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commentRes)
}

func (h *Handler) GetComment(c *gin.Context) {
	commentId := c.Query("id")

	comment, err := h.CommentService.Get(context.Background(), &pbc.Id{CommentId: commentId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *Handler) DeleteComment(c *gin.Context) {
	commentId := c.Query("id")

	_, err := h.CommentService.Delete(context.Background(), &pbc.Id{CommentId: commentId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment was deleted successfully",
	})
}

func (h *Handler) ListComments(c *gin.Context) {
	comments, err := h.CommentService.List(context.Background(), &pbc.GetListFilter{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}
