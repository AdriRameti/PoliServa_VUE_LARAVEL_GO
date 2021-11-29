package reservations

import (
	"fmt"
	"log"
	"net/http"
	"poliserva/Config"

	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	var reservations []ReservationModel

	if err := Config.DB.Find(&reservations).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	c.JSON(http.StatusOK, reservations)
}

func GetDateReservations(c *gin.Context){
	Date:= c.Query("Fini")
	Id_court:= c.Query("Id_Court")

	if len(Finit)<1 {
		log.Println("Fini is missing")
		return
	}
	
}
