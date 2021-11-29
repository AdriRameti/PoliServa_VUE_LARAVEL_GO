package reservations

import (
	"github.com/gin-gonic/gin"
)

func ReservationRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllReservations)
	router.GET("/datereservation", GetDateReservations)
}