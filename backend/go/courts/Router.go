package courts

import (
	"github.com/gin-gonic/gin"
)

func CourtsRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllCourts)
	router.GET("/sport", GetCourtsBySport)
	router.GET("/single", GetOneCourt)
	router.POST("/", CreateCourt)
	router.PUT("/", UpdateCourt)
	router.DELETE("/", DeleteCourt)
	
}