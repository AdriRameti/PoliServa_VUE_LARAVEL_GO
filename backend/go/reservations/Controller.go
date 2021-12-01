package reservations

import (
	"log"
	"poliserva/Common"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	var reservations []ReservationModel
	consult:=GetAllReservation(&reservations,c)
	serializer:= ReservationsSerializer{c, *consult}
	c.JSON(http.StatusCreated,serializer.Response())
}

func GetDateReservations(c *gin.Context){
	var reservations []ReservationModel
	Date:= c.Query("date")
	Hour:= c.Query("hour")
	if len(Date)<1||  len(Hour)<1{
		return
	}
	validDate:= common.ValidateDate(Date)
	validHour:= common.ValidateHour(Hour)
	if validDate == false || validHour == false{
		c.AbortWithStatus(http.StatusNotFound);
	}
	MyArray:= []string {Date,Hour}
	consult:=GetDateOneReservation(MyArray,&reservations,c)
	serializer:= ReservationsSerializer{c, *consult}
	c.JSON(http.StatusCreated, serializer.Response())
}
func CreateReservations(c *gin.Context){
	reservationModelValidation := NewReservationModelValidator()
	if err := reservationModelValidation.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Error Validation Reservation")
		return
	}
	CreateNewReservation(&reservationModelValidation.reservationModel,c)
	serializer:= ReservationSerializer{c, reservationModelValidation.reservationModel}
	c.JSON(http.StatusCreated, serializer.Response())

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
