package handlers

import (
	"encoding/json"
	"exam/api-gateway/api-testing/storage"
	"exam/api-gateway/api-testing/storage/kv"
	"fmt"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// User crud
func RegisterUser(c *gin.Context) {
	var newUser storage.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error(), "1")
		return
	}

	newUser.Id = uuid.NewString()
	newUser.Email = strings.ToLower(newUser.Email)
	err := newUser.Validate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error(), "2")
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error(), "3")
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error(), "4")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "a verification code was sent to your email, please check it.",
	})
}

func Verify(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Correct code",
	})
}

func CreateUser(c *gin.Context) {
	var newUser storage.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	userString, err := kv.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.UserRequest
	if err := json.Unmarshal([]byte(userString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// pp.Println(resp, "get user")
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if err := kv.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func ListUsers(c *gin.Context) {
	usersStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*storage.UserRequest
	for _, userString := range usersStrings {
		var user storage.UserRequest
		if err := json.Unmarshal([]byte(userString), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}

// Post crud
func CreatePost(c *gin.Context) {
	var newProduct storage.Post
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postJson, err := json.Marshal(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newProduct.Id), string(postJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func GetPost(c *gin.Context) {
	productID := c.Query("id")
	postString, err := kv.Get(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.Post
	if err := json.Unmarshal([]byte(postString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeletePost(c *gin.Context) {
	productId := c.Query("id")
	if err := kv.Delete(productId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post was deleted successfully",
	})
}

func ListPosts(c *gin.Context) {
	productsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var posts []*storage.Post
	for _, productString := range productsStrings {
		var post storage.Post
		if err := json.Unmarshal([]byte(productString), &post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		posts = append(posts, &post)
	}

	c.JSON(http.StatusOK, posts)
}

// Comment crud
func CreateComment(c *gin.Context) {
	var newComment storage.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentJson, err := json.Marshal(newComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newComment.Id), string(commentJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newComment)
}

func GetComment(c *gin.Context) {
	commentID := c.Query("id")
	postString, err := kv.Get(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.Comment
	if err := json.Unmarshal([]byte(postString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteComment(c *gin.Context) {
	commentId := c.Query("id")
	if err := kv.Delete(commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment was deleted successfully",
	})
}

func ListComments(c *gin.Context) {
	commentStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var comments []*storage.Comment
	for _, commentString := range commentStrings {
		var comment storage.Comment
		if err := json.Unmarshal([]byte(commentString), &comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		comments = append(comments, &comment)
	}

	c.JSON(http.StatusOK, comments)
}
