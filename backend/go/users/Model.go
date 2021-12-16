package users

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
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

func GetUserWithRoleDB(uuid string, c *gin.Context) UserModel {
	
	var user UserModel

	if err := Config.DB.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return user
	} else {
		return user
	}

}

func (b *UserModel) TableName() string {
	return "users"
}