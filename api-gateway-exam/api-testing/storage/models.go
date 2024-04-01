package storage

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type UserRequest struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
	Bio          string `json:"bio"`
	Website      string `json:"website"`
}

type Post struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	Content  string `json:"content"`
	Likes    int64  `json:"likes"`
	Dislikes int64  `json:"dislikes"`
	Views    int64  `json:"views"`
	Category string `json:"category"`
}

type Comment struct {
	Id      string `json:"id"`
	PostId  string `json:"post_id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}

// User info validation
func (u *UserRequest) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email))
}

type Message struct {
	Message string `json:"message"`
}
