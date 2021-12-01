package reservations

import (
	// "fmt"
	"log"
	// "net/http"
	// "poliserva/Config"
	// "strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	var reservations []ReservationModel

	GetAllReservation(&reservations,c)
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
	// var reservations ReservationModel
	reservationModelValidation := NewReservationModelValidator()
	
	if err := reservationModelValidation.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Error Validation Reservation")
		return
	}
	// c.BindJSON(&reservations)
	CreateNewReservation(&reservationModelValidation.reservationModel,c)

}

func UpdateReservations(c *gin.Context){
	var reservations ReservationModel
	id := c.Params.ByName("id")
	if len(id)<1{
		log.Println("Id is missing")
		return
	}
	UpdateReservation(&reservations,id,c)
}

func DeleteReservations(c *gin.Context){
	var reservations ReservationModel
	id := c.Params.ByName("id")
	if len(id)<1{
		log.Println("Id is missing")
		return
	}
	DeleteReservation(&reservations,id,c)
}
