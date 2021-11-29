package reservations

import (
	"fmt"
	"log"
	"net/http"
	"poliserva/Config"
	// "strconv"

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
	var reservations []ReservationModel
	Date:= c.Query("date")
	Id_court:= c.Query("court")
	Hour:= c.Query("hour")
	if len(Date)<1{
		log.Println("Date is missing")
		return
	}
	if len(Id_court)<1{
		log.Println("Court is missing")
		return
	}
	if len(Hour)<1{
		log.Println("Hour is missing")
		return
	}
	MyArray:= []string {Date,Id_court,Hour}
	GetDateOneReservation(MyArray,&reservations,c)
}
func CreateReservations(c *gin.Context){
	var reservations ReservationModel

	c.BindJSON(&reservations)

	CreateNewReservation(&reservations,c)

}
