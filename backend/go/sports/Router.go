package sports

import (
	"github.com/gin-gonic/gin"
)

func SportsRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllSports)
	
}