package courts

import (
	"github.com/gin-gonic/gin"
)

func CourtsRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllCourts)
	router.GET("/datereservation", GetDateReservations)
	router.GET("/sport", GetCourtsBySport)
	router.GET("/carousel", GetCourtsCarousel)
	router.GET("/single", GetOneCourt)
	router.POST("/", CreateCourt)
	router.PUT("/", UpdateCourt)
	router.DELETE("/", DeleteCourt)
	
}