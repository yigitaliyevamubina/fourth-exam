package tests

import (
	"encoding/json"
	pbc "exam/api-gateway/genproto/comment_service"
	pbp "exam/api-gateway/genproto/post_service"
	pbu "exam/api-gateway/genproto/user_service"

	"fmt"

	"exam/api-gateway/mock/handlers"

	"exam/api-gateway/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Message struct {
	Message string `json:"message"`
}

func TestMockMicroserviceApi(t *testing.T) {
	buffer, err := mock.OpenFile("user.json")
	require.NoError(t, err)

	h := handlers.NewHandler(&mock.UserServiceClient{}, &mock.PostServiceClient{}, &mock.CommentServiceClient{})

	//User CRUD Test

	// User Create
	req := mock.NewRequest(http.MethodPost, "/users/create", buffer)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/users/create", h.CreateUser)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var user pbu.User
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	// fmt.Println(user, "-----------")
	require.Equal(t, user.Email, "testemail@gmail.com")
	require.Equal(t, user.FirstName, "Test FirstName 1")
	require.Equal(t, user.LastName, "Test Lastname 1")
	require.Equal(t, user.Bio, "Test bio 1")
	require.Equal(t, user.Website, "Test website 1")
	require.Equal(t, user.Username, "Test Username 1")
	require.Equal(t, user.RefreshToken, "refresh token test")
	require.Equal(t, user.Password, "**kkw##knrtest")
	require.Equal(t, user.RefreshToken, "refresh token test")

	// User Get
	getReq := mock.NewRequest(http.MethodGet, "/users/get", buffer)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", h.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp pbu.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Email, getUserResp.Email)
	assert.Equal(t, user.Username, getUserResp.Username)
	assert.Equal(t, user.Bio, getUserResp.Bio)
	assert.Equal(t, user.Website, getUserResp.Website)

	// User List
	listReq := mock.NewRequest(http.MethodGet, "/users", buffer)
	listRes := httptest.NewRecorder()

	r.GET("/users", h.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	var listUsersResp pbu.Users
	require.NoError(t, json.Unmarshal(bodyBytes, &listUsersResp))
	require.Equal(t, listUsersResp.Count, int64(3))

	// User Delete
	delReq := mock.NewRequest(http.MethodDelete, "/users/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	r.DELETE("/users/delete", h.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var resMessage Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resMessage))
	require.Equal(t, "user was deleted successfully", resMessage.Message)

	// Check field
	body := pbu.CheckFieldReq{
		Field: "email",
		Value: "testemail@gmail.com",
	}
	buffer, err = json.Marshal(body)
	assert.NoError(t, err)
	checkFieldReq := mock.NewRequest(http.MethodPost, "/users/checkfield", buffer)
	checkFieldRes := httptest.NewRecorder()
	r.POST("/users/checkfield", h.CheckField)
	r.ServeHTTP(checkFieldRes, checkFieldReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, checkFieldRes.Code)
	var message Message
	bodyBytes, err = io.ReadAll(checkFieldRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user exists", message.Message)

	// Update refresh token
	bodyReq := pbu.UpdateRefreshReq{
		UserId:       user.Id,
		RefreshToken: "refresh token test",
	}
	buffer, err = json.Marshal(bodyReq)
	assert.NoError(t, err)
	updateRefreshReq := mock.NewRequest(http.MethodPost, "/users/update/refreshtoken", buffer)
	updateRefreshRes := httptest.NewRecorder()
	r.POST("/users/update/refreshtoken", h.UpdateRefresh)
	r.ServeHTTP(updateRefreshRes, updateRefreshReq)
	var updateRefresh pbu.User
	require.NoError(t, json.Unmarshal(updateRefreshRes.Body.Bytes(), &updateRefresh))
	assert.Equal(t, user.Id, updateRefresh.Id)
	assert.Equal(t, user.FirstName, updateRefresh.FirstName)
	assert.Equal(t, user.LastName, updateRefresh.LastName)
	assert.Equal(t, user.Email, updateRefresh.Email)

	//Post CRUD Test
	buffer, err = mock.OpenFile("post.json")
	require.NoError(t, err)

	// Post create
	reqPost := mock.NewRequest(http.MethodPost, "/posts/create", buffer)
	resPost := httptest.NewRecorder()
	r = gin.Default()
	r.POST("/posts/create", h.CreatePost)
	r.ServeHTTP(resPost, reqPost)
	assert.Equal(t, http.StatusOK, resPost.Code)

	var post pbp.Post
	require.NoError(t, json.Unmarshal(resPost.Body.Bytes(), &post))
	require.Equal(t, post.Id, "e292ca9d-d202-4aa2-a7de-487158b02dd4")
	require.Equal(t, post.UserId, "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf")
	require.Equal(t, post.Content, "Test Content 1")
	require.Equal(t, post.Likes, int64(10))
	require.Equal(t, post.Dislikes, int64(10))
	require.Equal(t, post.Views, int64(10))
	require.Equal(t, post.Category, "Test Category 1")

	// Post Get
	getReq = mock.NewRequest(http.MethodGet, "/posts/get", buffer)
	q = getReq.URL.Query()
	q.Add("id", cast.ToString(post.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/posts/get", h.GetPost)
	r.ServeHTTP(getRes, getReq)
	fmt.Println(getRes.Body)
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

	// Post List
	listReq = mock.NewRequest(http.MethodGet, "/posts", buffer)
	listRes = httptest.NewRecorder()

	r.GET("/posts", h.ListPosts)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	var listPostsResp pbp.Posts
	require.NoError(t, json.Unmarshal(bodyBytes, &listPostsResp))
	require.Equal(t, listPostsResp.Count, int64(3))

	// Product Delete
	delReq = mock.NewRequest(http.MethodDelete, "/posts/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", cast.ToString(post.Id))
	delReq.URL.RawQuery = q.Encode()
	delRes = httptest.NewRecorder()
	r.DELETE("/posts/delete", h.DeletePost)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respMessage Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &respMessage))
	require.Equal(t, "post was deleted successfully", respMessage.Message)

	//Comment CRUD Test
	buffer, err = mock.OpenFile("comment.json")
	require.NoError(t, err)

	// Comment create
	reqComment := mock.NewRequest(http.MethodPost, "/comments/create", buffer)
	resComment := httptest.NewRecorder()
	r = gin.Default()
	r.POST("/comments/create", h.CreateComment)
	r.ServeHTTP(resComment, reqComment)
	assert.Equal(t, http.StatusOK, resComment.Code)

	var comment pbc.Comment
	require.NoError(t, json.Unmarshal(resComment.Body.Bytes(), &comment))
	require.Equal(t, comment.Id, "")
	require.Equal(t, comment.PostId, "e292ca9d-d202-4aa2-a7de-487158b02dd4")
	require.Equal(t, comment.UserId, "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf")
	require.Equal(t, comment.Content, "Test Content 1")

	// Comment Get
	getReq = mock.NewRequest(http.MethodGet, "/comments/get", buffer)
	q = getReq.URL.Query()
	q.Add("id", cast.ToString(post.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/comments/get", h.GetComment)
	r.ServeHTTP(getRes, getReq)
	fmt.Println(getRes.Body)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getCommentResp pbc.Comment
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getCommentResp))
	fmt.Println(getCommentResp)
	require.Equal(t, comment.Id, getCommentResp.Id)
	require.Equal(t, comment.PostId, getCommentResp.PostId)
	require.Equal(t, comment.UserId, getCommentResp.UserId)
	require.Equal(t, comment.Content, getCommentResp.Content)

	// Comment List
	listReq = mock.NewRequest(http.MethodGet, "/comments", buffer)
	listRes = httptest.NewRecorder()

	r.GET("/comments", h.ListComments)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	var listCommentsResp pbc.Comments
	require.NoError(t, json.Unmarshal(bodyBytes, &listCommentsResp))
	require.Equal(t, listCommentsResp.Count, int64(3))

	// Comment Delete
	delReq = mock.NewRequest(http.MethodDelete, "/comments/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", cast.ToString(post.Id))
	delReq.URL.RawQuery = q.Encode()
	delRes = httptest.NewRecorder()
	r.DELETE("/comments/delete", h.DeleteComment)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respCommentMessage Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &respCommentMessage))
	require.Equal(t, "comment was deleted successfully", respCommentMessage.Message)
}
