package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRouting(router *gin.RouterGroup) {

	router.POST("/getrole", GetUserWithRole)
}