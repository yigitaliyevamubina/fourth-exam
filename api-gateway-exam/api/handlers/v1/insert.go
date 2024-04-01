package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"
	"exam/api-gateway/microservice-testing/models.go"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Insert data to db(mongo)
// @Router /v1/insert [get]
// @Security BearerAuth
// @Summary data insertion
// @Tags Insertion
// @Description Insert data to db(mongo)
// @Accept json
// @Produce json
// @Success 201 {object} models.Message
// @Failure 400 string Error models.Message
// @Failure 500 string Error models.Message
func (h *handlerV1) InsertToMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	// Mock users
	users := []*pbu.User{}
	for i := 0; i < 10; i++ {
		user := &pbu.User{
			Id:           uuid.New().String(),
			Username:     fmt.Sprintf("user%d", i+1),
			Email:        fmt.Sprintf("user%d@example.com", i+1),
			Password:     fmt.Sprintf("password%d", i+1),
			FirstName:    fmt.Sprintf("User%d_FirstName", i+1),
			LastName:     fmt.Sprintf("User%d_LastName", i+1),
			Bio:          fmt.Sprintf("Bio for User%d", i+1),
			Website:      fmt.Sprintf("www.user%dwebsite.com", i+1),
			CreatedAt:    time.Now().Format(time.RFC3339),
			UpdatedAt:    time.Now().Format(time.RFC3339),
			IsActive:     true,
			RefreshToken: fmt.Sprintf("token%d", i+1),
		}
		_, err := h.serviceManager.UserService().Create(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Message{
				Message: err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	// Mock posts
	posts := []*pbp.Post{}
	for i := 0; i < 10; i++ {
		post := &pbp.Post{
			Id:        uuid.New().String(),
			UserId:    users[i].Id,
			Content:   fmt.Sprintf("Content of Post%d", i+1),
			Title:     fmt.Sprintf("Title of Post%d", i+1),
			Likes:     int64(i + 1),
			Dislikes:  int64(i),
			Views:     int64((i + 1) * 100),
			Category:  fmt.Sprintf("Category%d", i+1),
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
		_, err := h.serviceManager.PostService().Create(ctx, post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Message{
				Message: err.Error(),
			})
			return
		}
		posts = append(posts, post)
	}

	// Mock comments
	for i := 0; i < 10; i++ {
		comment := &pbc.Comment{
			Id:        uuid.New().String(),
			PostId:    posts[i].Id,
			UserId:    users[i].Id,
			Content:   fmt.Sprintf("Comment Content for Post%d", i+1),
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
		_, err := h.serviceManager.CommentService().Create(ctx, comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Message{
				Message: err.Error(),
			})
			return
		}
	}


	c.JSON(http.StatusOK, models.Message{
		Message: "Successfully inserted data to users, posts and comments. Now you can try other api",
	})
}
