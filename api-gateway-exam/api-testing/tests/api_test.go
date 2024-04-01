package tests

import (
	"encoding/json"
	"exam/api-gateway/api-testing/handlers"
	"exam/api-gateway/api-testing/storage"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err := OpenFile("user.json")
	// fmt.Println(err, " +++++  user json")
	require.NoError(t, err)

	// User Create
	req := NewRequest(http.MethodPost, "/users/create", buffer)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/users/create", handlers.CreateUser)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var user storage.UserRequest
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	fmt.Println(user, "-----------")
	require.Equal(t, user.Email, "testemail@gmail.com")
	require.Equal(t, user.FirstName, "TestFirstName")
	require.Equal(t, user.LastName, "TestLastname")
	require.Equal(t, user.Password, "**kkw##knrtest")
	require.Equal(t, user.Bio, "Test bio 1")
	require.Equal(t, user.Website, "Test website 1")

	require.NotNil(t, user.Id)

	// User Get
	getReq := NewRequest(http.MethodGet, "/users/get", buffer)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", handlers.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp storage.UserRequest
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Email, getUserResp.Email)

	// User List
	listReq := NewRequest(http.MethodGet, "/users", buffer)
	listRes := httptest.NewRecorder()

	r.GET("/users", handlers.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// User Delete
	delReq := NewRequest(http.MethodDelete, "/users/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	r.DELETE("/users/delete", handlers.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var message storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user was deleted successfully", message.Message)

	// // User Register
	regReq := NewRequest(http.MethodPost, "/users/register", buffer)
	regRes := httptest.NewRecorder()
	r.POST("/users/register", handlers.RegisterUser)
	r.ServeHTTP(regRes, regReq)
	// fmt.Println(string(buffer))
	assert.Equal(t, http.StatusOK, regRes.Code)
	var resp storage.Message
	bodyBytes, err = io.ReadAll(regRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resp))
	require.NotNil(t, resp.Message)
	require.Equal(t, "a verification code was sent to your email, please check it.", resp.Message)

	// User Verify
	uri := fmt.Sprintf("/users/verify/%s", "12345")
	verReq := NewRequest(http.MethodGet, uri, buffer)
	verRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/verify/:code", handlers.Verify)
	r.ServeHTTP(verRes, verReq)
	assert.Equal(t, http.StatusOK, verRes.Code)
	var response *storage.Message
	bodyBytes, err = io.ReadAll(verRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &response))
	require.Equal(t, "Correct code", response.Message)

	//User Verify with incorrect code
	incorrectURI := fmt.Sprintf("/users/verify/%s", "11111")
	incorrectVerReq := NewRequest(http.MethodGet, incorrectURI, buffer)
	incorrectVerRes := httptest.NewRecorder()
	r.ServeHTTP(incorrectVerRes, incorrectVerReq)
	assert.Equal(t, http.StatusBadRequest, incorrectVerRes.Code)
	var incorrectResponse storage.Message
	bodyBytes, err = io.ReadAll(incorrectVerRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &incorrectResponse))
	require.Equal(t, "Incorrect code", incorrectResponse.Message)

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err = OpenFile("post.json")
	require.NoError(t, err)

	// Post Create
	req = NewRequest(http.MethodPost, "/posts/create", buffer)
	// fmt.Println(string(buffer))
	res = httptest.NewRecorder()
	r.POST("/posts/create", handlers.CreatePost)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var post storage.Post
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &post))
	require.Equal(t, "Test Content 1", post.Content)
	require.Equal(t, "Test Category 1", post.Category)
	require.Equal(t, int64(10), post.Likes)
	require.Equal(t, int64(10), post.Dislikes)

	// Post Get
	getReq = NewRequest(http.MethodGet, "/posts/get", buffer)
	q = getReq.URL.Query()
	q.Add("id", string(post.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/posts/get", handlers.GetPost)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)

	var getPost storage.Post
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getPost))
	// pp.Println(string(bodyBytes))
	require.Equal(t, post.Category, getPost.Category)
	require.Equal(t, post.Content, getPost.Content)

	// Product List
	listReq = NewRequest(http.MethodGet, "/posts", buffer)
	listRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/posts", handlers.ListPosts)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	// pp.Println(string(bodyBytes))

	// Product Delete
	delReq = NewRequest(http.MethodDelete, "/posts/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", string(post.Id))
	delReq.URL.RawQuery = q.Encode()
	r.DELETE("/posts/delete", handlers.DeletePost)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var postMessage storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &postMessage))
	require.Equal(t, "post was deleted successfully", postMessage.Message)

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err = OpenFile("comment.json")
	require.NoError(t, err)

	// Comment Create
	req = NewRequest(http.MethodPost, "/comments/create", buffer)
	// fmt.Println(string(buffer))
	res = httptest.NewRecorder()
	r.POST("/comments/create", handlers.CreateComment)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var comment storage.Comment
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &comment))
	require.Equal(t, "Test Content 1", comment.Content)
	require.Equal(t, "e292ca9d-d202-4aa2-a7de-487158b02dd4", comment.PostId)

	// Comment Get
	getReq = NewRequest(http.MethodGet, "/comments/get", buffer)
	q = getReq.URL.Query()
	q.Add("id", string(comment.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/comments/get", handlers.GetComment)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)

	var getComment storage.Comment
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getComment))
	// pp.Println(string(bodyBytes))
	require.Equal(t, comment.Content, getComment.Content)
	require.Equal(t, comment.PostId, getComment.PostId)
	require.Equal(t, comment.UserId, getComment.UserId)

	// Comment List
	listReq = NewRequest(http.MethodGet, "/comments", buffer)
	listRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/comments", handlers.ListComments)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	// pp.Println(string(bodyBytes))

	// Comment Delete
	delReq = NewRequest(http.MethodDelete, "/comments/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", string(comment.Id))
	delReq.URL.RawQuery = q.Encode()
	r.DELETE("comments/delete", handlers.DeleteComment)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var commentMessage storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &commentMessage))
	require.Equal(t, "comment was deleted successfully", commentMessage.Message)
}
