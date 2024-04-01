package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type UserRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Biography string `json:"biography"`
	Website   string `json:"website"`
	IsActive  bool   `json:"isActive"`
}

type RegisterUserModel struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Biography string `json:"biography"`
	Website   string `json:"website"`
	IsActive  bool   `json:"is_active"`
	Code      string `json:"code"`
}

// User info validation
func (u *UserRequest) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 15), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
		validation.Field(&u.LastName, validation.Required, validation.Length(5, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
	)
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UserModel struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Biography    string `json:"biography"`
	Website      string `json:"website"`
	IsActive     bool   `json:"is_active"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ListUsers struct {
	Count int64               `json:"count"`
	Users []*UserWithProducts `json:"users"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenUpdateReq struct {
	RefreshToken string `json:"refresh_token"`
}

type Status struct {
	Success bool `string:"status"`
}

type UserWithProducts struct {
	User  UserModel           `json:"user"`
	Posts []*PostWithComments `json:"posts"`
}
