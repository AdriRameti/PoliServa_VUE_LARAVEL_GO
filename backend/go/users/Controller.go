package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	var users []UserModel

	users = GetAllUsersRP(c)

	serializer := UsersSerializer{c, users}

	c.JSON(http.StatusOK, serializer.Response())
}

func CreateUser(c *gin.Context) {

	userModelValidator := NewUserValidator()

	if err, strErr := userModelValidator.Bind(c); (err != nil || strErr == "already used") || (err != nil && strErr == "can't bind") {
		c.JSON(200, strErr)
		return
	}

	user, errC := CreateUserRP(&userModelValidator.userModel, c)

	if errC != nil {
		c.JSON(http.StatusUnprocessableEntity, errC)
		return
	}

	serializer := UserSerializer{c, user}

	c.JSON(http.StatusOK, serializer.Response())
}

func UpdateUser(c *gin.Context) {

	var uuid string

	uuid = c.Query("uuid")

	userModelValidator := NewUserValidator()

	if err, strErr := userModelValidator.Bind(c); (err != nil || strErr == "already used") || (err != nil && strErr == "can't bind") {
		c.JSON(http.StatusUnprocessableEntity, strErr)
		return
	}

	user, errU := UpdateUserRP(uuid, &userModelValidator.userModel, c)

	if errU != nil {
		c.JSON(http.StatusUnprocessableEntity, errU)
		return
	}

	serializer := UserSerializer{c, user}

	c.JSON(http.StatusOK, serializer.Response())
}

func DeleteUser(c *gin.Context) {
	var uuid string

	uuid = c.Query("uuid")

	DeleteUserRP(uuid, c)
}

func GetUserWithRole(c *gin.Context) {
	var uuid string
	var user UserModel

	uuid = c.GetHeader("uuid")

	user = GetUserWithRoleRP(uuid, c)

	serializer := UserSerializer{c, user}

	c.JSON(http.StatusOK, serializer.Response())
}