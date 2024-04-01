package microservicetesting

import (
	"bytes"
	"encoding/json"
	"exam/api-gateway/config"
	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"
	"exam/api-gateway/microservice-testing/handlers"
	"exam/api-gateway/microservice-testing/models.go"

	"github.com/brianvoe/gofakeit/v6"

	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RunApiTest(t *testing.T) {
	// New IServiceManager
	cfg := config.Load()
	service, err := services.NewServiceManager(&cfg)
	require.NoError(t, err)
	sm := handlers.NewMockServiceManager(service, logger.New("", ""))

	//tests

	//Create User
	id := uuid.New().String()
	email := fmt.Sprintf(gofakeit.Noun() + "@gmail.com")
	username := gofakeit.Username()
	user := &pbu.User{
		Id:           id,
		FirstName:    "Test FirstName 1",
		LastName:     "Test Lastname 1",
		Username:     username,
		Bio:          "Test bio 1",
		Website:      "Test website 1",
		Email:        email,
		Password:     "**kkw##knrtest",
		RefreshToken: "refresh token test",
	}
	payloadBytes, err := json.Marshal(user)
	require.NoError(t, err)
	//response
	r := gin.Default()
	r.POST("/users/create", sm.CreateUser)
	req, err := http.NewRequest(http.MethodPost, "/users/create", bytes.NewReader(payloadBytes))
	require.NoError(t, err)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var respUser *pbu.User
	assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &respUser))
	require.Equal(t, user.Email, email)
	require.Equal(t, user.FirstName, "Test FirstName 1")
	require.Equal(t, user.LastName, "Test Lastname 1")
	require.Equal(t, user.Bio, "Test bio 1")
	require.Equal(t, user.Website, "Test website 1")
	require.Equal(t, user.Username, username)
	require.Equal(t, user.RefreshToken, "refresh token test")
	require.Equal(t, user.Password, "**kkw##knrtest")
	require.Equal(t, user.RefreshToken, "refresh token test")

	// Get User by id
	getReq, err := http.NewRequest(http.MethodGet, "/users/get", nil)
	require.NoError(t, err)
	q := getReq.URL.Query()
	q.Add("id", id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.GET("/users/get", sm.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp *pbu.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Email, getUserResp.Email)
	assert.Equal(t, user.Username, getUserResp.Username)
	assert.Equal(t, user.Bio, getUserResp.Bio)
	assert.Equal(t, user.Website, getUserResp.Website)

	user.FirstName = "Updated first name"
	user.LastName = "Updated last name"
	payloadBytes, err = json.Marshal(user)
	require.NoError(t, err)

	// Update User
	updateReq, err := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer(payloadBytes))
	require.NoError(t, err)
	updateRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.PUT("/users/update", sm.UpdateUser)
	r.ServeHTTP(updateRes, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var updatedUser *pbu.User
	assert.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &updatedUser))
	assert.Equal(t, user.FirstName, updatedUser.FirstName)
	assert.Equal(t, user.LastName, updatedUser.LastName)

	// Post
	post := &pbp.Post{
		Id:       uuid.New().String(),
		UserId:   user.Id,
		Content:  "Test Content 1",
		Likes:    10,
		Dislikes: 10,
		Views:    10,
		Category: "Test Category 1",
	}
	// Create post
	postBytes, err := json.Marshal(post)
	require.NoError(t, err)
	//response
	r = gin.Default()
	r.POST("/posts/create", sm.CreatePost)
	createProdReq, err := http.NewRequest(http.MethodPost, "/posts/create", bytes.NewReader(postBytes))
	require.NoError(t, err)
	createProdRes := httptest.NewRecorder()
	r.ServeHTTP(createProdRes, createProdReq)
	assert.Equal(t, http.StatusOK, createProdRes.Code)
	var respPost *pbp.Post
	assert.NoError(t, json.Unmarshal(createProdRes.Body.Bytes(), &respPost))
	require.Equal(t, post.UserId, user.Id)
	require.Equal(t, post.Content, "Test Content 1")
	require.Equal(t, post.Likes, int64(10))
	require.Equal(t, post.Dislikes, int64(10))
	require.Equal(t, post.Views, int64(10))
	require.Equal(t, post.Category, "Test Category 1")

	// Get Post by id
	getReq, err = http.NewRequest(http.MethodGet, "/posts/get", nil)
	require.NoError(t, err)
	q = getReq.URL.Query()
	q.Add("id", post.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.GET("/posts/get", sm.GetPost)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getPostResp pbp.Post
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getPostResp))
	fmt.Println(getPostResp)
	require.Equal(t, post.Id, getPostResp.Id)
	require.Equal(t, post.UserId, getPostResp.UserId)
	require.Equal(t, post.Content, getPostResp.Content)
	require.Equal(t, post.Likes, getPostResp.Likes)
	require.Equal(t, post.Dislikes, getPostResp.Dislikes)
	require.Equal(t, post.Views, getPostResp.Views)
	require.Equal(t, post.Category, getPostResp.Category)

	post.Content = "Updated content"
	payloadBytes, err = json.Marshal(post)
	require.NoError(t, err)

	// Update Post
	updateReq, err = http.NewRequest(http.MethodPut, "/posts/update", bytes.NewBuffer(payloadBytes))
	require.NoError(t, err)
	updateRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.PUT("/posts/update", sm.UpdatePost)
	r.ServeHTTP(updateRes, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var updatedPost *pbp.Post
	assert.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &updatedPost))
	assert.Equal(t, post.Content, updatedPost.Content)

	// Comment
	id = uuid.New().String()
	commentReq := &pbc.Comment{
		Id:      id,
		PostId:  post.Id,
		UserId:  user.Id,
		Content: "Test Content 1",
	}
	fmt.Println(commentReq, "   comment")
	// Create comment
	commentBytes, err := json.Marshal(commentReq)
	require.NoError(t, err)
	//response
	r = gin.Default()
	r.POST("/comments/create", sm.CreateComment)
	reqComment, err := http.NewRequest(http.MethodPost, "/comments/create", bytes.NewReader(commentBytes))
	require.NoError(t, err)
	resComment := httptest.NewRecorder()
	r.ServeHTTP(resComment, reqComment)
	assert.Equal(t, http.StatusOK, createProdRes.Code)
	var comment pbc.Comment
	require.NoError(t, json.Unmarshal(resComment.Body.Bytes(), &comment))
	require.Equal(t, commentReq.Id, id)
	require.Equal(t, commentReq.PostId, post.Id)
	require.Equal(t, commentReq.UserId, user.Id)
	require.Equal(t, commentReq.Content, "Test Content 1")

	// Get Comment by id
	getReq, err = http.NewRequest(http.MethodGet, "/comments/get", nil)
	require.NoError(t, err)
	q = getReq.URL.Query()
	q.Add("id", id)
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.GET("/comments/get", sm.GetComment)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getCommentResp pbc.Comment
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getCommentResp))
	fmt.Println(getCommentResp)
	require.Equal(t, comment.Id, getCommentResp.Id)
	require.Equal(t, comment.PostId, getCommentResp.PostId)
	require.Equal(t, comment.UserId, getCommentResp.UserId)
	require.Equal(t, comment.Content, getCommentResp.Content)

	comment.Content = "Updated content"
	payloadBytes, err = json.Marshal(comment)
	require.NoError(t, err)

	// Update Comment
	updateReq, err = http.NewRequest(http.MethodPut, "/comments/update", bytes.NewBuffer(payloadBytes))
	require.NoError(t, err)
	updateRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.PUT("/comments/update", sm.UpdateComment)
	r.ServeHTTP(updateRes, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var updatedComment *pbc.Comment
	assert.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &updatedComment))
	assert.Equal(t, comment.Content, updatedPost.Content)

	// Delete Comment by id
	delReq, err := http.NewRequest(http.MethodDelete, "/comments/delete", nil)
	require.NoError(t, err)
	q = delReq.URL.Query()
	q.Add("id", comment.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.DELETE("/comments/delete", sm.DeleteComment)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code) 
	var delCommMessage models.Message
	bodyBytes, err := io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &delCommMessage))
	require.Equal(t, "comment was deleted", delCommMessage.Message)

	// Delete Post by id
	delReq, err = http.NewRequest(http.MethodDelete, "/posts/delete", nil)
	require.NoError(t, err)
	q = delReq.URL.Query()
	q.Add("id", respPost.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.DELETE("/posts/delete", sm.DeletePost)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var delProdMessage models.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &delProdMessage))
	require.Equal(t, "post was deleted", delProdMessage.Message)

	// Delete User by id
	delReq, err = http.NewRequest(http.MethodDelete, "/users/delete", nil)
	require.NoError(t, err)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.DELETE("/users/delete", sm.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var message models.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user was deleted", message.Message)
}
