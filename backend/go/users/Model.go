package users

import (
	// "poliserva/Config"
	// "fmt"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

type UserModel struct {
	ID uint `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Surnames string `json:"surnames"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
	Img string `json:"img"`
}

// func (uuid string, c *gin.Context) GetUserWithRoleDB() {

// }

func (b *UserModel) TableName() string {
	return "users"
}