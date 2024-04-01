package v1

import (
	"context"
	"exam/api-gateway/api/handlers/models"
	pb "exam/api-gateway/genproto/comment_service"
	"exam/api-gateway/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateComment
// @Router /v1/comment/create [post]
// @Security BearerAuth
// @Summary create comment
// @Tags Comment
// @Description Create a new comment with provided details
// @Accept json
// @Produce json
// @Param CommentDetails body models.Comment true "Create comment"
// @Success 201 {object} models.Comment
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		body       models.Comment
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

	resp, err := h.serviceManager.CommentService().Create(ctx, &pb.Comment{
		Id: uuid.New().String(),
		PostId: body.PostID,
		UserId: body.UserID,
		Content: body.Content,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot create comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update Comment
// @Router /v1/comment/update/{id} [put]
// @Security BearerAuth
// @Summary update comment
// @Tags Comment
// @Description Update comment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param CommentInfo body models.Comment true "Update Comment"
// @Success 201 {object} models.Comment
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var (
		body        models.Comment
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

	response, err := h.serviceManager.CommentService().Update(ctx, &pb.Comment{
		Id: id,
		PostId: body.PostID,
		UserId: body.UserID,
		Content: body.Content,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot update comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Comment By Id
// @Router /v1/comment/{id} [get]
// @Security BearerAuth
// @Summary get comment by id
// @Tags Comment
// @Description Get comment
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 201 {object} models.Comment
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) GetCommentById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.CommentService().Get(ctx, &pb.Id{CommentId: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})
		h.log.Error("cannot get comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Comment
// @Router /v1/comment/delete/{id} [delete]
// @Security BearerAuth
// @Summary delete comment
// @Tags Comment
// @Description Delete comment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Status
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) DeleteComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	id := c.Param("id")

	_, err := h.serviceManager.CommentService().Delete(ctx, &pb.Id{CommentId: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot delete comment", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.Status{
		Success: true,
	})
}

// Get All Comments
// @Router /v1/comments/{page}/{limit} [get]
// @Summary get all comments
// @Tags Comment
// @Description get all comments
// @Accept json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Param orderBy query string false "orderBy" Enums(content, created_at, updated_at) "Order by"
// @Success 201 {object} models.ListComments
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) ListComments(c *gin.Context) {
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

	response, err := h.serviceManager.CommentService().List(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot list comments", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get All Comments by user id
// @Router /v1/comments/{page}/{limit}/{user_id} [get]
// @Summary get all comments
// @Tags Comment
// @Description get all comments by user id
// @Accept json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Param orderBy query string false "orderBy" Enums(content, created_at, updated_at) "Order by"
// @Param user_id path string true "user_id"
// @Success 201 {object} models.ListComments
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) ListCommentsByUserId(c *gin.Context) {
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

	response, err := h.serviceManager.CommentService().List(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot list comments by user id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get All Comments by post id
// @Router /v1/get/comments/{page}/{limit}/{post_id} [get]
// @Summary get all comments
// @Tags Comment
// @Description get all comments by post id
// @Accept json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Param orderBy query string false "orderBy" Enums(content, created_at, updated_at) "Order by"
// @Param post_id path string true "post_id"
// @Success 201 {object} models.ListComments
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
func (h *handlerV1) ListCommentsByPostId(c *gin.Context) {
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

	postId := c.Param("post_id")
	req.PostId = postId

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

	response, err := h.serviceManager.CommentService().List(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternalServerError,
			Message: err.Error(),
		})

		h.log.Error("cannot list comments by post id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
