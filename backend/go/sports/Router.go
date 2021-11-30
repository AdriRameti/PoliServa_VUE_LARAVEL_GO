package sports

import (
	"github.com/gin-gonic/gin"
)

func SportsRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllSports)
	router.GET("/single", GetOneSport)
	router.POST("/", CreateSport)
	router.PUT("/", UpdateSport)
	router.DELETE("/", DeleteSport)
	
}