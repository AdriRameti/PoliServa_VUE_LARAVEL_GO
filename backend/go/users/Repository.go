package users

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAllUsersRP (c *gin.Context) []UserModel
	GetUserWithRoleRP (uuid string, c *gin.Context) UserModel
	CreateUserRP (user *UserModel, c *gin.Context) (userR UserModel, errR error)
	UpdateUserRP (uuid string, user *UserModel, c *gin.Context) (userR UserModel, errR error)
	DeleteUserRP (uuid string)
	CheckMail (user *UserModelValidator, c *gin.Context) bool
}

func GetAllUsersRP(c *gin.Context) []UserModel {
	var users[] UserModel

	if err := Config.DB.Find(&users).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return users
	} else {
		return users
	}
}

func GetUserWithRoleRP(uuid string, c *gin.Context) UserModel {
	
	var user UserModel

	if err := Config.DB.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return user
	} else {
		return user
	}

}

func CreateUserRP(user *UserModel, c *gin.Context) (userR UserModel, errR error) {

	fmt.Println("ME CAGUE EN DEU")

	var userRet UserModel

	if err := Config.DB.Create(user).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		
		return userRet, err
	} else {
		fmt.Println("hola si create")
		Config.DB.Where("uuid = ?", user.UUID).Find(&userRet)
		return userRet, err
	}
}

func UpdateUserRP(uuid string, user *UserModel, c *gin.Context) (userR UserModel, errR error) {

	var userF UserModel

	if errF := Config.DB.Where("uuid = ?", uuid).Find(&userF).Error; errF == nil {

		userF.Name = user.Name
		userF.Surnames = user.Surnames
		userF.Mail = user.Mail
		userF.Pass = user.Pass
		userF.Img = user.Img
		userF.Role = user.Role
		userF.Is_blocked = user.Is_blocked

		Config.DB.Save(&userF)

		return userF, errF
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return userF, errF
	}
}

func DeleteUserRP(uuid string, c *gin.Context) {

	var user UserModel

	if errF := Config.DB.Where("uuid = ?", uuid).Delete(&user).Error; errF == nil {
		c.JSON(http.StatusOK, "Se ha eliminado con éxito")
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

}


func CheckMail(user *UserModelValidator, c *gin.Context) bool {
	var Newuser UserModel

	if err := Config.DB.Where("mail = ?", user.Mail).Find(&Newuser).Error; err != nil {
		return false
	} else {

		if (user.UUID == Newuser.UUID) {
			return false
		} else {
			return true
		}
	}
 
}