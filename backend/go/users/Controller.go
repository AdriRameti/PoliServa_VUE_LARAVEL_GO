package users

import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
)

func GetUserWithRole(c *gin.Context) {
	var uuid string

	uuid = c.GetHeader("uuid")

	c.JSON(http.StatusOK, uuid)

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