package reservations

import (
	"github.com/gin-gonic/gin"
)

func ReservationRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllReservations)
	router.GET("/datereservation", GetDateReservations)
	router.GET("/userreservation",GetUserReservation)
	router.GET("/usersportreservation",GetUserSportReservation)
	router.POST("/",CreateReservations)
	router.PUT("/:id",UpdateReservations)
	router.DELETE("/:id",DeleteReservations)
}