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