package v1

import (
	"context"
	"exam/api-gateway/api/handlers/models"
	pb "exam/api-gateway/genproto/post_service"
	"exam/api-gateway/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost
// @Router /v1/post/create [post]
// @Security BearerAuth
// @Summary create post
// @Tags Post
// @Description Insert a new post with provided details
// @Accept json
// @Produce json
// @Param PostDetails body models.Post true "Create post"
// @Success 201 {object} models.Post
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body       models.Post
		jspMarshal protojson.MarshalOptions
	)
	jspMarshal.UseProtoNames = true

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	resp, err := h.serviceManager.PostService().Create(ctx, &pb.Post{
		Id:       uuid.New().String(),
		UserId:   body.UserID,
		Content:  body.Content,
		Title:    body.Title,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
		Category: body.Category,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot create post", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update Post
// @Router /v1/post/update/{id} [put]
// @Security BearerAuth
// @Summary update post
// @Tags Post
// @Description Update post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param PostInfo body models.Post true "Update Post"
// @Success 201 {object} models.Post
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        models.Post
		jspbMarshal protojson.MarshalOptions
	)
	id := c.Param("id")

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot bind json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.PostService().Update(ctx, &pb.Post{
		Id:       id,
		UserId:   body.UserID,
		Content:  body.Content,
		Title:    body.Title,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
		Category: body.Category,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot update post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Like Post
// @Router /v1/post/like [put]
// @Security BearerAuth
// @Summary like post
// @Tags Post
// @Description Like post
// @Accept json
// @Produce json
// @Param post_id body models.PostReq true "Like Post"
// @Success 201 {object} models.Post
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) LikePost(c *gin.Context) {
	var (
		body        models.PostReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot bind json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	post, err := h.serviceManager.PostService().Get(context.Background(), &pb.Id{PostId: body.PostId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot get post by id", logger.Error(err))
		return
	}

	post.Likes += 1

	response, err := h.serviceManager.PostService().Update(ctx, &pb.Post{
		Id:       body.PostId,
		UserId:   post.UserId,
		Content:  post.Content,
		Title:    post.Title,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Views:    post.Views,
		Category: post.Category,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot update post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Like Post
// @Router /v1/post/dislike [put]
// @Security BearerAuth
// @Summary dislike post
// @Tags Post
// @Description Dislike post
// @Accept json
// @Produce json
// @Param post_id body models.PostReq true "Dislike Post"
// @Success 201 {object} models.Post
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) DislikePost(c *gin.Context) {
	var (
		body        models.PostReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot bind json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	post, err := h.serviceManager.PostService().Get(context.Background(), &pb.Id{PostId: body.PostId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot get post by id", logger.Error(err))
		return
	}

	post.Dislikes -= 1

	response, err := h.serviceManager.PostService().Update(ctx, &pb.Post{
		Id:       body.PostId,
		UserId:   post.UserId,
		Content:  post.Content,
		Title:    post.Title,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Views:    post.Views,
		Category: post.Category,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot update post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Post By Id
// @Router /v1/post/{id} [get]
// @Security BearerAuth
// @Summary get post by id
// @Tags Post
// @Description Get post
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 201 {object} models.Post
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) GetPostById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.PostService().Get(ctx, &pb.Id{PostId: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot get post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Post
// @Router /v1/post/delete/{id} [delete]
// @Security BearerAuth
// @Summary delete post
// @Tags Post
// @Description Delete post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Status
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	id := c.Param("id")

	_, err := h.serviceManager.PostService().Delete(ctx, &pb.Id{PostId: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot delete post", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.Status{
		Success: true,
	})
}

// Get All Posts
// @Router /v1/posts/{page}/{limit} [get]
// @Summary get all posts
// @Tags Post
// @Description get all posts
// @Accept json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Param orderBy query string false "orderBy" Enums(content, title, category) "Order by"
// @Success 201 {object} models.ListPosts
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) ListPosts(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		req pb.GetListFilter
	)
	orderBy := c.Query("orderBy")
	if orderBy != "" {
		req.OrderBy = orderBy
	} else {
		req.OrderBy = "created_at"
	}

	page := c.Param("page")
	pageToInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot parse page query param", logger.Error(err))
		return
	}

	limit := c.Param("limit")
	LimitToInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot parse limit query param", logger.Error(err))
		return
	}

	req.Page = int64(pageToInt)
	req.Limit = int64(LimitToInt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.PostService().List(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot list posts", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get All Posts by user id 
// @Router /v1/posts/{page}/{limit}/{user_id} [get]
// @Summary get all posts
// @Tags Post
// @Description get all posts by user id
// @Accept json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Param orderBy query string false "orderBy" Enums(content, title, category, created_at, updated_at) "Order by"
// @Param user_id path string true "user_id"
// @Success 201 {object} models.ListPosts
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) ListPostsByUserId(c *gin.Context) { 
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		req pb.GetListFilter
	)
	orderBy := c.Query("orderBy")
	if orderBy != "" {
		req.OrderBy = orderBy
	} else {
		req.OrderBy = "created_at"
	}
	
	userId := c.Param("user_id")
	req.UserId = userId

	page := c.Param("page")
	pageToInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot parse page query param", logger.Error(err))
		return
	}

	limit := c.Param("limit")
	LimitToInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: err.Error(),
		})
		h.log.Error("cannot parse limit query param", logger.Error(err))
		return
	}

	req.Page = int64(pageToInt)
	req.Limit = int64(LimitToInt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.PostService().List(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot list posts by user id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
