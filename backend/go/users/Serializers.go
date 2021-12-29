package users

import (
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID uint `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Surnames string `json:"surnames"`
	Mail string `json:"mail"`
	Img string `json:"img"`
	Google2fa_secret string `json:"google2fa_secret"`
	Role string `json:"role"`
	Is_blocked bool `json:"is_blocked"`
}

type UserSerializer struct {
	C *gin.Context
	user UserModel
}

type UsersSerializer struct {
	C *gin.Context
	users []UserModel
}

func (us *UserSerializer) Response() UserResponse {
	response := UserResponse {
		ID: us.user.ID,
		UUID: us.user.UUID,
		Name: us.user.Name,
		Surnames: us.user.Surnames,
		Mail: us.user.Mail,
		Img: us.user.Img,
		Google2fa_secret: us.user.Google2fa_secret,
		Role: us.user.Role,
		Is_blocked: us.user.Is_blocked,
	}

	return response
}

func (us *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}

	for _, user := range us.users {
		serializer := UserSerializer{us.C, user}
		response = append(response, serializer.Response())
	}

	return response
}