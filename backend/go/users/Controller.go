package users

import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
)

func GetUserWithRole(c *gin.Context) {
	var uuid string
	var user UserModel

	uuid = c.GetHeader("uuid")

	user = GetUserWithRoleDB(uuid, c)

	serializer := UserSerializer{c, user}

	c.JSON(http.StatusOK, serializer.Response())

	// userWithRoleModelValidator := NewUserWithRoleValidator()

	// if err := userWithRoleModelValidator.Bind(c); err != nil {
	// 	// return "cannot bind uuid"
	// 	// fmt.Println("wawa")
	// 	c.JSON(http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, userWithRoleModelValidator.userWithRoleModel)

	// return userWithRoleModelValidator.userWithRoleModel;
}