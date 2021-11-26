package reservations

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"poliserva/Config"
)

func GetAllReservations(c *gin.Context) {
	var reservations []ReservationModel

	if err := Config.DB.Find(&reservations).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	c.JSON(http.StatusOK, reservations)
}